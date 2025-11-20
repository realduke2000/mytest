import time
import sys
import re
import inspect
from selenium import webdriver
from selenium.webdriver.common.by import By
from selenium.webdriver.chrome.options import Options as ChromeOptions
from selenium.webdriver.support.ui import WebDriverWait
from selenium.webdriver.support import expected_conditions as EC
from selenium.common.exceptions import (
    TimeoutException,
    ElementClickInterceptedException
)

# ===== 全局配置 =====
g_target_date = "2025-11-30"         # ⚠️ 请改成真实可选的日期
#g_agency = "agency-lgu"             # "agency-skt" | "agency-kt" | "agency-lgu" | "agency-and"
g_agency = "01"                      #  01（SKT）、02（KT）、03（LG U+）
g_user_name = "김분조"
g_user_phone = "010-5101-5251"       # 会自动清洗为纯数字
# 仅用于 SMS 验证
g_ssn6 = "580721"                    # 주민등록번호前6位：YYMMDD
g_ssn1 = "2"                         # 주민등록번호第7位（性别/世纪位）：1/2/3/4 等

enable_debugger = False
SMS_Verification = True              # True 走 SMS 验证；False 保持 PASS 验证

# ===== 工具函数 =====
STATUS_STARTED = "Started"
STATUS_SUCCESS = "Success"
STATUS_FAILED = "Failed"
STATUS_SKIPPED = "Skipped"

def dom_ready(driver):
    return driver.execute_script("return document.readyState") == "complete"

def where_am_i(driver):
    """根据已知关键元素判断当前KCB页面所处阶段"""
    try:
        if driver.find_elements(By.CSS_SELECTOR, '.mobileCoCheck[value]'):
            return "carrier_select"   # 通信社选择页
        if driver.find_elements(By.CSS_SELECTOR, '.mobileCertMethodCheck[value]') and driver.find_elements(By.ID, 'btnMobileCertStart'):
            return "method_select"    # 认证方式选择页（SMS/PASS/QR）
        if driver.find_elements(By.ID, "authForm") and driver.find_elements(By.ID, "nm"):
            return "sms_input"        # SMS 输入页（姓名/生日/手机号...）
        return "unknown"
    except Exception:
        return "unknown"

def func_name(level=0):
    """level=0 当前函数名，1 调用者，2 调用者的调用者"""
    return inspect.stack()[level].function

def log_step(id, description, status):
    print(f"\n Step {id}: {description} - {status}")
    time.sleep(0.2)

def safe_click_by_id(driver, element_id, timeout=15):
    """等待可点击 + 滚动 + 遮挡兜底（JS）"""
    wait = WebDriverWait(driver, timeout)
    el = wait.until(EC.element_to_be_clickable((By.ID, element_id)))
    driver.execute_script("arguments[0].scrollIntoView({block:'center'});", el)
    time.sleep(0.1)
    try:
        el.click()
    except ElementClickInterceptedException:
        driver.execute_script(f"document.getElementById('{element_id}').click();")

def switch_to_new_window(driver, before_handles, timeout=10):
    """等待新窗口出现并切换"""
    wait = WebDriverWait(driver, timeout)
    wait.until(lambda d: len(d.window_handles) > len(before_handles))
    new_handle = (set(driver.window_handles) - set(before_handles)).pop()
    driver.switch_to.window(new_handle)
    return new_handle

class Progress:
    def __init__(self, totalSteps=10):
        self.step = 0
        self.totalSteps = totalSteps
    
    def nextStep(self):
        self.step += 1
        return self.step


def OpenAppointmentPage(driver, stepID, timeout=30):
    desc = ' '.join(re.findall(r'[A-Z]?[a-z]+|[A-Z]+(?![a-z])',func_name()))
    log_step(stepID, desc, STATUS_STARTED)

    if enable_debugger:
        driver.get("about:blank")
        time.sleep(10)

    wait = WebDriverWait(driver, timeout=timeout, poll_frequency=1)
    driver.get("https://www.chronodigmwatch.co.kr/rolex/contact-seoul/appointment")
    wait.until(lambda d: d.execute_script("return document.readyState") == "complete")
    log_step(stepID, desc, STATUS_SUCCESS)

def AcceptCookie(driver, stepID, timeout=30):
    # Step 1.1: 接受 cookie
    desc = ' '.join(re.findall(r'[A-Z]?[a-z]+|[A-Z]+(?![a-z])',func_name()))
    try:
        log_step(stepID, desc, STATUS_STARTED)
        wait = WebDriverWait(driver, timeout=timeout, poll_frequency=1)
        log_step("Step 1.1: click cookie popup if exists")
        cookie_btn = wait.until(EC.element_to_be_clickable(
            (By.XPATH, '/html/body/div[1]/div[1]/div/div/button[2]')
        ))
        cookie_btn.click()
        log_step(stepID, desc, STATUS_SUCCESS)
    except Exception:
        log_step(stepID, desc, STATUS_SKIPPED)

def ClickAppointmentService(driver, stepID, timeout=30):
    # Step 2: Click appointment service “서비스 관련 시계 접수 및 수령”
    desc = ' '.join(re.findall(r'[A-Z]?[a-z]+|[A-Z]+(?![a-z])',func_name()))

    log_step(stepID, desc, STATUS_STARTED)
    wait = WebDriverWait(driver, timeout=timeout, poll_frequency=1)
    elem = wait.until(EC.presence_of_element_located((By.XPATH, '//*[@id="fappointment"]/div[1]/div/div/a[1]')))
    driver.execute_script("arguments[0].scrollIntoView({behavior: 'smooth', block: 'center'});", elem)
    wait.until(EC.element_to_be_clickable((By.XPATH, '//*[@id="fappointment"]/div[1]/div/div/a[1]'))).click()
    log_step(stepID, desc, STATUS_SUCCESS)
    
    if enable_debugger:
        # update js get_datetime_list()
        # set data.popup as ""
        time.sleep(30)
    

# ===== 主流程 =====
def run_chronodigm_appointment_v8(target_date, agency, user_name, user_phone):
    chrome_opts = ChromeOptions()
    # 如需无痕/无头可自行开启
    # chrome_opts.add_argument("--incognito")
    # chrome_opts.add_argument("--headless=new")

    driver = webdriver.Chrome(options=chrome_opts)
    wait = WebDriverWait(driver, timeout=60, poll_frequency=1)

    try:
        progress = Progress()
        OpenAppointmentPage(driver, progress.nextStep())
        # 记录原窗口句柄，后续实名成功后切回
        original_handle = driver.current_window_handle
        stepID += 1
        AcceptCookie(driver, stepID)
        stepID += 1
        ClickAppointmentService(driver,stepID)

        # Step 3.1: 点击 “동의합니다”
        log_step("Step 3.1: click agree button")
        wait.until(EC.element_to_be_clickable((By.XPATH, '//*[@id="fappointment"]/div[2]/footer/button'))).click()

        # Step 3.2: 检查是否有“结束”弹窗
        log_step("Step 3.2: check end_popup")
        try:
            wait.until(EC.presence_of_element_located((By.ID, "end_popup")))
            print("appointment is not open to book")
            raise Exception("appointment is not open to book")
        except TimeoutException:
            print("no end_popup, proceed")

        # Step 4: 选择预约日期
        log_step(f"Step 4: select appointment date {target_date}")
        date_xpath = f'//li[@data-date="{target_date}"]'
        wait.until(EC.element_to_be_clickable((By.XPATH, date_xpath))).click()

        # Step 5: 选择最早时间
        log_step("Step 5: select earliest timeslot")
        slot_container_xpath = f'//div[@data-date="{target_date}" and contains(@class, "time-slot") and contains(@style, "display: block")]'
        wait.until(EC.visibility_of_element_located((By.XPATH, slot_container_xpath)))
        time_items = driver.find_elements(By.XPATH, f'{slot_container_xpath}//li[@data-time and not(contains(@class, "off"))]')
        if not time_items:
            raise Exception("no available timeslot")
        first_time_item = time_items[0]
        driver.execute_script("arguments[0].scrollIntoView({behavior: 'smooth', block: 'center'});", first_time_item)
        time.sleep(0.3)
        driver.execute_script("arguments[0].click();", first_time_item)
        print("clicked time slot:", first_time_item.text)
        wait.until(lambda d: "active" in first_time_item.get_attribute("class"))
        print("selected time slot successfully")

        while True:
            try:
                # Step 6: 点击“다음”按钮，打开验证窗口
                log_step("Step 6: click next to open safe.ok-name")
                before_handles = driver.window_handles[:]
                print("original windows:", before_handles)

                next_btn = wait.until(EC.element_to_be_clickable((By.XPATH, '//button[contains(text(),"다음")]')))
                driver.execute_script("arguments[0].scrollIntoView({behavior: 'smooth', block: 'center'});", next_btn)
                time.sleep(0.3)
                next_btn.click()
                print("clicked next -> open safe.ok-name")

                # Step 7: 切换到新窗口
                log_step("Step 7: switch to safe.ok-name window")
                try:
                    new_handle = switch_to_new_window(driver, before_handles, timeout=20)
                    print("switched to KCB window:", new_handle)

                    # 等待页面加载完成
                    WebDriverWait(driver, 20).until(dom_ready)

                    # 等待窗口中的某个已知阶段出现
                    stage = WebDriverWait(driver, 20).until(lambda d: (s := where_am_i(d)) != "unknown" and s)
                    print("KCB landing stage:", stage)
                except Exception as e:
                    print("Step 7 exception:", repr(e))
                    print("current page URL:", driver.current_url)
                    print("current page title:", driver.title)
                    raise

                # Step 8: 选择运营商 + 选择认证方式（SMS验证）
                log_step("Step 8: choose agency and SMS verification method")
                try:
                    # 选择运营商按钮
                    carrier_btn = wait.until(EC.element_to_be_clickable((By.CSS_SELECTOR, f'.mobileCoCheck[value="{agency}"]')))
                    driver.execute_script("arguments[0].scrollIntoView({block:'center'});", carrier_btn)
                    time.sleep(0.2)
                    carrier_btn.click()
                    print("clicked agency:", agency)

                    # 等待认证方式选项加载（SMS/QR/PASS）
                    wait.until(EC.presence_of_element_located((By.CSS_SELECTOR, '.cert_list')))
                    print("certification methods loaded")

                    # 选择SMS认证按钮
                    sms_btn = wait.until(EC.element_to_be_clickable((By.CSS_SELECTOR, '.mobileCertMethodCheck[value="SMS"]')))
                    driver.execute_script("arguments[0].scrollIntoView({block:'center'});", sms_btn)
                    time.sleep(0.3)
                    sms_btn.click()
                    print("clicked SMS verification button")

                    # 同意使用协议
                    agree_checkbox = wait.until(EC.presence_of_element_located((By.ID, "mobileCertAgree")))
                    if not agree_checkbox.is_selected():
                        agree_checkbox.click()
                        print("clicked agree to terms checkbox")

                    # 点击“다음”按钮提交
                    next_btn = wait.until(EC.element_to_be_clickable((By.ID, "btnMobileCertStart")))
                    driver.execute_script("arguments[0].scrollIntoView({block:'center'});", next_btn)
                    time.sleep(0.3)
                    next_btn.click()
                    print("clicked next to submit verification method")
                    
                except Exception as e:
                    print("Step 8 exception:", repr(e))
                    screenshot_path = f"/mnt/data/step8_failed_{time.strftime('%Y%m%d-%H%M%S')}.png"
                    driver.save_screenshot(screenshot_path)
                    print(f"screenshot saved to {screenshot_path}")
                    raise

                # Step 9: 分支 — SMS 验证 或 PASS 验证
                if SMS_Verification:
                    # ===== SMS 验证 =====
                    log_step("Step 9: SMS - switch tab and fill SMS form")

                    # 9.1 填写姓名
                    name_el = wait.until(EC.presence_of_element_located((By.ID, "nm")))
                    name_el.clear()
                    name_el.send_keys(g_user_name)
                    print("filled name")
                    
                    # 点击 "다음" 按钮继续
                    next_btn = wait.until(
                        EC.element_to_be_clickable((By.CLASS_NAME, "btn_pass.btnUserName"))
                    )
                    driver.execute_script("arguments[0].scrollIntoView({behavior: 'smooth', block: 'center'});", next_btn)
                    time.sleep(0.3)
                    next_btn.click()  # 点击“다음”按钮
                    print("clicked '다음' button to proceed")

                    # 9.2 填写住址（6位出生年月日）+ 1位性别
                    ssn6_el = wait.until(EC.presence_of_element_located((By.ID, "ssn6")))
                    ssn6_el.clear()
                    ssn6_el.send_keys(g_ssn6[:6])
                    print("filled ssn6")

                    ssn1_el = wait.until(EC.presence_of_element_located((By.ID, "ssn1")))
                    ssn1_el.clear()
                    ssn1_el.send_keys(g_ssn1[:1])
                    print("filled ssn1")

                    # 9.3 填写手机号（纯数字11位）
                    phone_digits = re.sub(r"\D", "", g_user_phone)[:11]
                    phone_el = wait.until(EC.presence_of_element_located((By.ID, "mbphn_no")))
                    phone_el.clear()
                    phone_el.send_keys(phone_digits)
                    print("filled phone number:", phone_digits)

                    # 点击 "다음" 按钮继续提交
                    next_btn = wait.until(EC.element_to_be_clickable((By.CLASS_NAME, "btn_pass")))
                    driver.execute_script("arguments[0].scrollIntoView({behavior: 'smooth', block: 'center'});", next_btn)
                    time.sleep(0.3)
                    next_btn.click()  # 点击“다음”按钮提交表单
                    print("clicked '다음' button to submit form")

                    # 后续等待
                    ans = input("Continue to step 10 and finish appointment? ")
                    if ans.lower() in ["yes", "y"]:
                        if original_handle in driver.window_handles:
                            driver.switch_to.window(original_handle)
                            print("switched back to original window")
                            continue
                else:
                    # ===== PASS 验证（保持原逻辑）=====
                    log_step("Step 9: PASS - fill pass form")

                    # 按你原逻辑：nm / mbphn_no / captchaCode
                    name_input = wait.until(EC.presence_of_element_located((By.ID, "nm")))
                    name_input.clear()
                    name_input.send_keys(g_user_name)
                    print("filled name")

                    phone_digits = re.sub(r"\D", "", g_user_phone)[:11]
                    phone_input = wait.until(EC.presence_of_element_located((By.ID, "mbphn_no")))
                    phone_input.clear()
                    phone_input.send_keys(phone_digits)
                    print("filled phone:", phone_digits)

                    # 验证码
                    wait.until(EC.presence_of_element_located((By.ID, "botDetectCaptcha_CaptchaImage")))
                    captcha_img = driver.find_element(By.ID, "botDetectCaptcha_CaptchaImage")
                    print("captcha url:", captcha_img.get_attribute("src"))
                    captcha_code = input("please input captcha (PASS): ").strip()

                    captcha_input = wait.until(EC.presence_of_element_located((By.ID, "captchaCode")))
                    captcha_input.clear()
                    captcha_input.send_keys(captcha_code)
                    print("filled captchaCode")

                # Step 10: 提交并切回原窗口
                log_step("Step 10: submit verification and switch back")
                try:
                    confirm_btn = wait.until(EC.element_to_be_clickable((By.ID, "btnSubmit")))
                    driver.execute_script("arguments[0].scrollIntoView({block:'center'});", confirm_btn)
                    time.sleep(0.2)
                    try:
                        confirm_btn.click()
                    except ElementClickInterceptedException:
                        driver.execute_script("document.getElementById('btnSubmit').click();")
                    print("clicked btnSubmit")

                    # 给服务端处理时间
                    time.sleep(2)

                    # 如需回主窗口（常见流程是回到预约窗口继续）
                    if original_handle in driver.window_handles:
                        driver.switch_to.window(original_handle)
                        print("switched back to original window")
                except Exception as e:
                    print("Step 10 exception:", repr(e))
                    raise

                # 后续你可以在这里继续预约流程的后半段…
                time.sleep(1)
            except Exception as e:
                ans=input("try again?")
                if ans == "yes" or ans == "y" or ans == "Y":
                    if original_handle in driver.window_handles:
                        driver.switch_to.window(original_handle)
                        print("switched back to original window")
                    continue
                else:
                    raise


    except Exception as e:
        print(f"\nuncaught exception: {e}")
        print(str(e))
    finally:
        driver.quit()

# ===== 启动入口 =====
if __name__ == "__main__":
    target_date = g_target_date
    if len(sys.argv) > 1 and sys.argv[1].strip():
        target_date = sys.argv[1].strip()
    run_chronodigm_appointment_v8(target_date, g_agency, g_user_name, g_user_phone)
