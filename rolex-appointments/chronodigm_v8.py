import time
from selenium import webdriver
from selenium.webdriver.common.by import By
from selenium.webdriver.firefox.options import Options
from selenium.webdriver.support.ui import WebDriverWait
from selenium.webdriver.support import expected_conditions as EC

# 📅 设置目标预约日期
target_date = "2025-07-30"  # 根据实际情况修改

def log_step(step):
    print(f"\n🔹 {step} - 开始")
    time.sleep(0.3)

def run_chronodigm_appointment_v8():
    options = Options()
    driver = webdriver.Firefox(options=options)
    wait = WebDriverWait(driver, 15)

    try:
        # Step 1: 打开页面
        log_step("Step 1: 打开预约页面")
        driver.get("https://www.chronodigmwatch.co.kr/rolex/contact-seoul/appointment")
        WebDriverWait(driver, 10).until(lambda d: d.execute_script("return document.readyState") == "complete")

        # Step 1.1: 接受 cookie
        try:
            log_step("Step 1.1: 尝试点击 cookie 弹窗")
            cookie_btn = WebDriverWait(driver, 3).until(EC.element_to_be_clickable(
                (By.XPATH, '/html/body/div[1]/div[1]/div/div/button[2]')
            ))
            cookie_btn.click()
            print("✅ 已点击 cookie 弹窗")
        except:
            print("⚠️ 未检测到 cookie 弹窗，继续")

        # Step 2: 点击 “서비스 관련 시계 접수 및 수령”
        log_step("Step 2: 点击服务类型按钮")
        wait.until(EC.element_to_be_clickable((
            By.XPATH, '/html/body/div[2]/main/section[1]/div/div/div[2]/form[1]/div[1]/div/div/a[2]'
        ))).click()

        # Step 3: 点击 “동의합니다”
        log_step("Step 3: 点击同意按钮")
        wait.until(EC.element_to_be_clickable((
            By.XPATH, '/html/body/div[2]/main/section[1]/div/div/div[2]/form[1]/div[2]/footer/button'
        ))).click()

        # Step 4: 选择预约日期
        log_step(f"Step 4: 选择预约日期 {target_date}")
        date_xpath = f'//li[@data-date="{target_date}"]'
        wait.until(EC.element_to_be_clickable((By.XPATH, date_xpath))).click()

        # Step 5: 选择最早时间
        log_step("Step 5: 等待并点击该日最早可用时间")
        slot_container_xpath = f'//div[@data-date="{target_date}" and contains(@class, "time-slot") and contains(@style, "display: block")]'
        WebDriverWait(driver, 7).until(EC.visibility_of_element_located((By.XPATH, slot_container_xpath)))
        time_items = driver.find_elements(By.XPATH, f'{slot_container_xpath}//li[@data-time and not(contains(@class, "off"))]')
        if not time_items:
            raise Exception("❌ 没有可用的预约时间！")
        first_time_item = time_items[0]
        driver.execute_script("arguments[0].scrollIntoView({behavior: 'smooth', block: 'center'});", first_time_item)
        time.sleep(0.5)
        driver.execute_script("arguments[0].click();", first_time_item)
        print("✅ 点击了时间：", first_time_item.text)
        WebDriverWait(driver, 5).until(lambda d: "active" in first_time_item.get_attribute("class"))
        print("✅ 时间选择成功，状态变为 active")

        # Step 6: 点击 “다음”
        log_step("Step 6: 点击 다음")

        # ✅ 应该在点击之前记录窗口句柄
        before_handles = driver.window_handles
        print("🪟 Step 6 前窗口句柄:", before_handles)

        next_btn = wait.until(EC.element_to_be_clickable((By.XPATH, '//button[contains(text(), "다음")]')))
        driver.execute_script("arguments[0].scrollIntoView({behavior: 'smooth', block: 'center'});", next_btn)
        time.sleep(0.5)
        next_btn.click()
        print("✅ 点击了 ‘다음’ 以打开实名认证窗口")


        # Step 7: 切换到实名认证窗口
        log_step("Step 7: 切换到实名认证窗口")

        try:
            # 等待新窗口弹出（窗口数量增加）
            WebDriverWait(driver, 10).until(lambda d: len(d.window_handles) > len(before_handles))
            after_handles = driver.window_handles
            print("🪟 当前所有窗口（Step 7）:", after_handles)

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


        # Step 8: 在新窗口中选择通信社 agency-kt
        log_step("Step 8: 在新窗口中选择通信社 agency-kt")
        try:
            wait.until(EC.presence_of_element_located((By.ID, 'agency-kt')))
            kt_btn = driver.find_element(By.ID, 'agency-kt')
            print("📌 找到元素:", kt_btn.get_attribute("outerHTML"))
            driver.execute_script("arguments[0].scrollIntoView();", kt_btn)
            time.sleep(0.5)
            kt_btn.click()
            print("✅ 已点击 agency-kt")
        except Exception as e:
            print("❌ 未找到或点击运营商按钮:", e)
            raise e

    except Exception as e:
        print(f"\n❌ 出错: {e}")
    finally:
        driver.quit()

run_chronodigm_appointment_v8()
