import numpy as np
import matplotlib.pyplot as plt
from sklearn.linear_model import Ridge
from sklearn.preprocessing import PolynomialFeatures
from sklearn.pipeline import make_pipeline

# ==== 1. 原始实验数据 ====
data_raw = np.array([
    [0.4, 0.4, 68.9],
    [0.6, 0.6, 70.5],   # 修正过的真实数据
    [0.8, 0.8, 66.7],
    [0.4, 1.2, 72.2],
    [0.6, 1.0, 70.0],
    [0.8, 0.8, 68.5],
    [1.0, 0.6, 68.1],
    [1.2, 0.4, 69.7],
    [1.4, 0.2, 68.9]
])

d = data_raw[:, 0]
a = data_raw[:, 1]
measured = data_raw[:, 2]
baseline = 75.0
y = (baseline - measured) / baseline * 100  # 隔音率%

# ==== 2. 拟合 Ridge(Quadratic) 模型 ====
ridge_model = make_pipeline(PolynomialFeatures(2, include_bias=True), Ridge(alpha=1.0))
ridge_model.fit(np.column_stack([d, a]), y)

# 获取系数和截距
ridge_coef = ridge_model.named_steps['ridge'].coef_
ridge_intercept = ridge_model.named_steps['ridge'].intercept_

# 打印模型函数
print("Ridge(Quadratic) 模型：")
print(f"y(d, a) = {ridge_intercept:.4f} + {ridge_coef[1]:.4f}*d + {ridge_coef[2]:.4f}*a + "
      f"{ridge_coef[3]:.4f}*d^2 + {ridge_coef[4]:.4f}*d*a + {ridge_coef[5]:.4f}*a^2")

# ==== 3. 限制 d 和 a 的范围（d ∈ [0.4, 1.4], a ∈ [0.2, 1.2]） ====
d_values = np.linspace(0.4, 1.4, 100)  # 生成 d 范围内的 100 个数据点
a_values = np.linspace(0.2, 1.2, 100)  # 生成 a 范围内的 100 个数据点

# 创建网格数据
D, A = np.meshgrid(d_values, a_values)
X_grid = np.column_stack([D.ravel(), A.ravel()])

# 预测网格数据的隔音率
y_pred = ridge_model.predict(X_grid).reshape(D.shape)

# ==== 4. 绘制函数图像（3D图） ====
fig = plt.figure(figsize=(8, 6))
ax = fig.add_subplot(111, projection='3d')

# 绘制3D曲面
surf = ax.plot_surface(D, A, y_pred, cmap='viridis', linewidth=0, antialiased=True, alpha=0.9)

# 设置图标标签
ax.set_xlabel('d (cm)')
ax.set_ylabel('a (cm)')
ax.set_zlabel('Insulation rate (%)')
ax.set_title('Ridge(Quadratic) Model - 3D Surface')

# 设置视角
ax.view_init(elev=25, azim=45)

# 显示颜色条
fig.colorbar(surf, ax=ax, shrink=0.6)

plt.tight_layout()
plt.show()

plt.savefig("insulation_fit.png", dpi=300)
