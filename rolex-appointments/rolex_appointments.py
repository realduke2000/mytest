import time
from selenium import webdriver
from selenium.webdriver.common.by import By
from selenium.webdriver.firefox.options import Options
from selenium.webdriver.support.ui import WebDriverWait
from selenium.webdriver.support import expected_conditions as EC

from selenium.common.exceptions import NoSuchElementException,TimeoutException

# 📅 设置目标预约日期
target_date = "2025-08-31"  # 根据实际情况修改
# 设置姓名和手机号变量
user_name = "홍길동"       # 替换为实际姓名
user_phone = "01012345678" # 替换为实际手机号

def log_step(step):
    print(f"\n {step} - Started")
    time.sleep(0.3)

def run_chronodigm_appointment_v8():
    options = Options()
    driver = webdriver.Firefox(options=options)
    wait = WebDriverWait(driver, 15)

    try:
        # Step 1: 打开页面
        log_step("Step 1: Open appointment page")
        
        # mock, open debugger
        driver.get("about:blank")
        time.sleep(10)

        driver.get("https://www.chronodigmwatch.co.kr/rolex/contact-seoul/appointment")
        WebDriverWait(driver, 10).until(lambda d: d.execute_script("return document.readyState") == "complete")

        # Step 1.1: 接受 cookie
        try:
            log_step("Step 1.1: try to click cookie pop-up")
            cookie_btn = WebDriverWait(driver, 3).until(EC.element_to_be_clickable(
                (By.XPATH, '/html/body/div[1]/div[1]/div/div/button[2]')
            ))
            cookie_btn.click()
            print("✅ clicked cookie pop-up")
        except:
            print("did not detect cookie pop-up, continue")

        # Step 2: 点击 “서비스 관련 시계 접수 및 수령”
        log_step("Step 2: click appointment services")
        elem = wait.until(EC.presence_of_element_located((By.XPATH, '//*[@id="fappointment"]/div[1]/div/div/a[1]')))

        # scroll
        driver.execute_script("arguments[0].scrollIntoView({behavior: 'smooth', block: 'center'});", elem)

        # click it
        wait.until(EC.element_to_be_clickable((
            By.XPATH, '//*[@id="fappointment"]/div[1]/div/div/a[1]'
        ))).click()

        # mock, add breakpoint
        time.sleep(10)

        # Step 3.1: 点击 “동의합니다”
        log_step("Step 3.1: click agree button")
        wait.until(EC.element_to_be_clickable((
            By.XPATH, '//*[@id="fappointment"]/div[2]/footer/button'
        ))).click()

        # mock, clear data.popup
        print("clearing data.popup")
        time.sleep(30)

        # Step 3.2 
        log_step("Step 3.2 check if error dialog pop-up")
        try:
            popup = wait.until(EC.presence_of_element_located((By.ID, "end_popup")))
            print("appointment is not open to book")
            raise Exception("appointment is not open to book")
        except TimeoutException:
            print("not end_popup element, try to book appointment")
        


        # Step 4: 选择预约日期
        log_step(f"Step 4: select appointment date {target_date}")
        
        # mock available date
        target_li = driver.find_element(By.CSS_SELECTOR, '#datetime_form li[data-date="%s"]' % target_date)
        driver.execute_script("arguments[0].classList.remove('off');", target_li)

        date_xpath = f'//li[@data-date="{target_date}"]'
        wait.until(EC.element_to_be_clickable((By.XPATH, date_xpath))).click()

        # Step 5: 选择最早时间
        log_step("Step 5: wait and select the earliest timeslot")
        
        # mock available timeslot
        time_slot = driver.find_element(By.CSS_SELECTOR, f'.time-slot[data-date="{target_date}"]')
        li_element = time_slot.find_element(By.CSS_SELECTOR, f'li[data-time="960"]')
        driver.execute_script("arguments[0].classList.remove('off');", li_element)

        slot_container_xpath = f'//div[@data-date="{target_date}" and contains(@class, "time-slot") and contains(@style, "display: block")]'
        WebDriverWait(driver, 7).until(EC.visibility_of_element_located((By.XPATH, slot_container_xpath)))
        time_items = driver.find_elements(By.XPATH, f'{slot_container_xpath}//li[@data-time and not(contains(@class, "off"))]')
        if not time_items:
            raise Exception("no available timeslot")
        first_time_item = time_items[0]
        driver.execute_script("arguments[0].scrollIntoView({behavior: 'smooth', block: 'center'});", first_time_item)
        time.sleep(0.5)
        driver.execute_script("arguments[0].click();", first_time_item)
        print("clicked time slot:", first_time_item.text)
        WebDriverWait(driver, 5).until(lambda d: "active" in first_time_item.get_attribute("class"))
        print("selected time slot successfully")

        # Step 6: click “다음 >”
        log_step("Step 6: click next")

        # ✅ 应该在点击之前记录窗口句柄
        before_handles = driver.window_handles
        print(" Step 6 original windows handler:", before_handles)

        next_btn = wait.until(EC.element_to_be_clickable((By.XPATH, '//button[contains(text(),"다음")]')))
        driver.execute_script("arguments[0].scrollIntoView({behavior: 'smooth', block: 'center'});", next_btn)
        time.sleep(0.5)
        next_btn.click()
        print("clicked confirm to open safe.ok-name")


        # Step 7: 切换到实名认证窗口
        log_step("Step 7: switching to safe.ok-name window")

        try:
            # 等待新窗口弹出（窗口数量增加）
            WebDriverWait(driver, 10).until(lambda d: len(d.window_handles) > len(before_handles))
            after_handles = driver.window_handles
            print("all current windows（Step 7）:", after_handles)

            # 获取新窗口句柄
            new_windows = list(set(after_handles) - set(before_handles))
            if not new_windows:
                raise Exception("❌ 没有检测到新窗口句柄")
            new_window = new_windows[0]
            print("🪟 新窗口句柄:", new_window)

            # 切换到新窗口
            driver.switch_to.window(new_window)
            print("🔁 已切换至实名认证窗口")

            # 等待实名认证页面加载
            WebDriverWait(driver, 10).until(
                EC.presence_of_element_located((By.ID, "agree_all"))
            )
            print("✅ 检测到实名认证页面的 전체 동의 checkbox")

        except Exception as e:
            print("❌ Step 7 出错：", repr(e))
            print("🌐 当前页面 URL（失败时）：", driver.current_url)
            print("🧾 当前页面标题（失败时）：", driver.title)


        # Step 8: 在新窗口中选择通信社 agency-kt 并提交认证
        log_step("Step 8: 在新窗口中选择通信社 agency-kt")

        try:
            # 1. 选择运营商
            kt_radio = WebDriverWait(driver, 10).until(
                EC.presence_of_element_located((By.ID, "agency-kt"))
            )
            print(f"📌 找到元素: {kt_radio}")
            driver.execute_script("arguments[0].scrollIntoView({behavior: 'smooth', block: 'center'});", kt_radio)
            time.sleep(0.5)
            driver.execute_script("arguments[0].click();", kt_radio)
            print("✅ 成功点击 agency-kt")

            # 2. 勾选 “전체 동의하기”
            agree_checkbox = WebDriverWait(driver, 10).until(
                EC.presence_of_element_located((By.ID, "agree_all"))
            )
            driver.execute_script("arguments[0].scrollIntoView({behavior: 'smooth', block: 'center'});", agree_checkbox)
            time.sleep(0.5)
            if not agree_checkbox.is_selected():
                driver.execute_script("arguments[0].click();", agree_checkbox)
                print("✅ 已勾选 전체 동의하기")

            # 3. 点击 “인증하기” 按钮
            pass_btn = WebDriverWait(driver, 15).until(
                EC.element_to_be_clickable((By.ID, "btnPass"))
            )
            driver.execute_script("arguments[0].scrollIntoView({behavior: 'smooth', block: 'center'});", pass_btn)
            time.sleep(0.5)
            driver.execute_script("arguments[0].click();", pass_btn)
            print("🎉 已点击 ‘인증하기’，进入下一步实名验证流程")

        except Exception as e:
            print("❌ Step 8 出错：", repr(e))
            screenshot_path = f"/mnt/data/step8_failed_{time.strftime('%Y%m%d-%H%M%S')}.png"
            driver.save_screenshot(screenshot_path)
            print(f"📸 已保存截图：{screenshot_path}")
            raise
        
        # Step 9: 填写实名验证表单
        log_step("Step 9: 填写实名验证表单")
        try:
            # 等待姓名输入框加载并输入
            name_input = wait.until(EC.presence_of_element_located((By.ID, "nm")))
            name_input.clear()
            name_input.send_keys(user_name)
            print("✅ 姓名已填写")

            # 等待手机号输入框加载并输入
            phone_input = wait.until(EC.presence_of_element_located((By.ID, "mbphn_no")))
            phone_input.clear()
            phone_input.send_keys(user_phone)
            print("✅ 手机号已填写")

            # 等待验证码图片加载完毕
            captcha_img = wait.until(EC.presence_of_element_located((By.ID, "botDetectCaptcha_CaptchaImage")))
            captcha_src = captcha_img.get_attribute("src")
            print("📷 验证码图片链接：", captcha_src)

            # 暂停以人工输入验证码（也可后续接入 OCR）
            captcha_code = input("🔐 请输入验证码（从图片识别）：")

            # 填写验证码
            captcha_input = driver.find_element(By.ID, "captchaCode")
            captcha_input.clear()
            captcha_input.send_keys(captcha_code)
            print("✅ 验证码已填写")

        except Exception as e:
            print("❌ Step 9 出错：", repr(e))
            raise

        
        # Step 10: 点击 확인 提交实名验证
        log_step("Step 10: 提交实名验证")

        try:
            # 等待确认按钮可点击
            confirm_button = wait.until(EC.element_to_be_clickable((By.ID, "btnSubmit")))
            driver.execute_script("arguments[0].scrollIntoView({behavior: 'smooth', block: 'center'});", confirm_button)
            time.sleep(0.5)
            confirm_button.click()
            print("✅ 已点击 확인 按钮提交实名验证")
        except Exception as e:
            print("❌ Step 10 出错：", repr(e))
            raise


        time.sleep(30)



    except Exception as e:
        print(f"\nuncaught exception: {e}")
        print(str(e))
    finally:
        driver.quit()

run_chronodigm_appointment_v8()
