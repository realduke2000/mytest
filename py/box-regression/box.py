import numpy as np
import matplotlib.pyplot as plt
from sklearn.linear_model import Ridge
from sklearn.preprocessing import PolynomialFeatures
from sklearn.pipeline import make_pipeline
from mpl_toolkits.mplot3d import Axes3D

# ======================
# 1. 数据准备
# ======================
d = np.array([0.4, 0.6, 0.8, 0.4, 0.6, 0.8, 1.0, 1.2, 1.4])  # d (音源到内层盒的距离)
a = np.array([0.4, 0.6, 0.8, 1.2, 1.0, 0.8, 0.6, 0.4, 0.2])  # a (内层盒到外层盒的距离)
measured = np.array([68.9, 70.5, 66.7, 72.2, 70.0, 68.5, 68.1, 69.7, 68.9])  # 测得噪音值（dB）

baseline = 75.0  # 基准噪音值 75 dB
y = (baseline - measured) / baseline * 100  # 转换成百分比（隔音率）

X = np.column_stack([d, a])  # 将 d 和 a 组合成特征矩阵

# ======================
# 2. 训练 Ridge(Quadratic) 模型
# ======================
ridge_model = make_pipeline(PolynomialFeatures(2, include_bias=True), Ridge(alpha=1.0))
ridge_model.fit(X, y)

# 获取系数和截距
ridge_coef = ridge_model.named_steps['ridge'].coef_
ridge_intercept = ridge_model.named_steps['ridge'].intercept_

# 打印模型函数
print("Ridge(Quadratic) 模型：")
print(f"y(d, a) = {ridge_intercept:.4f} + {ridge_coef[1]:.4f}*d + {ridge_coef[2]:.4f}*a + "
      f"{ridge_coef[3]:.4f}*d^2 + {ridge_coef[4]:.4f}*d*a + {ridge_coef[5]:.4f}*a^2")

# ======================
# 3. 生成网格数据用于绘图
# ======================
d_grid = np.linspace(d.min(), d.max(), 50)
a_grid = np.linspace(a.min(), a.max(), 50)
D, A = np.meshgrid(d_grid, a_grid)
X_grid = np.column_stack([D.ravel(), A.ravel()])

# 预测网格点的隔音率
Y_pred = ridge_model.predict(X_grid).reshape(D.shape)

# ======================
# 4. 绘制三维图像
# ======================
fig = plt.figure(figsize=(8, 6))
ax = fig.add_subplot(111, projection='3d')

# 绘制拟合的3D曲面
surf = ax.plot_surface(D, A, Y_pred, cmap='viridis', linewidth=0, antialiased=True, alpha=0.9)

# 设置图标标签
ax.set_xlabel('d (cm)')
ax.set_ylabel('a (cm)')
ax.set_zlabel('Insulation rate (%)')
ax.set_title('Ridge(Quadratic) Model - 3D Surface')

# 设置视角
ax.view_init(elev=15, azim=55)

# 显示颜色条
fig.colorbar(surf, ax=ax, shrink=0.6)

plt.tight_layout()

plt.savefig("insulation_fit.png", dpi=300)
