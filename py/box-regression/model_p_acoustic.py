# -*- coding: utf-8 -*-
"""
P = f(d, h, a, n)
二次多项式特征 + 线性回归（最早版本的多元拟合）
- 读取你给出的表 5.1.1 / 5.1.2 / 5.2 / 5.4 的数据（已内置）
- 计算 CV 与综合性能 P（按你给的权重口径）
- 用 PolynomialFeatures(2) + LinearRegression 拟合
- 打印训练 R^2、系数表、并提供预测函数
- 含三张可视化（d-a-P 曲面、h-P 曲线、n-P 曲线）
"""

import numpy as np
import pandas as pd
import matplotlib.pyplot as plt
from sklearn.preprocessing import PolynomialFeatures
from sklearn.linear_model import LinearRegression

# -----------------------------
# 1) 录入数据（来自你提供的图片表格）
#   说明：
#   - 单层：a=0
#   - 双层（表5.2）：a=d（按图说明）
#   - 未指明 n 的，按常规实验视作 n=1
# -----------------------------
rows = []

# 表 5.1.1  单层，h=0.3，n=1，a=0，随 d 变化
data_511 = [
    # d, [f1000,f2000,f3000,f4000], f_t
    (0.4, [6.00, 5.60, 5.07, 4.67], 5.33),
    (0.6, [6.93, 6.67, 6.27, 5.73], 6.40),
    (0.8, [8.00, 7.73, 7.33, 6.80], 7.47),
    (1.0, [8.67, 8.40, 8.00, 7.47], 8.13),
    (1.2, [9.33, 8.93, 8.67, 8.13], 8.76),
    (1.4, [9.87, 9.60, 9.20, 8.67], 9.33),
    (1.6, [10.53, 10.27, 9.87, 9.33], 9.99),
    (1.8, [6.27, 6.00, 5.60, 5.20], 5.77),
    (2.2, [7.87, 7.60, 7.20, 6.80], 7.37),
    (2.6, [8.27, 8.00, 7.73, 7.33], 7.83),
    (3.0, [7.07, 6.80, 6.53, 6.13], 6.63),
    (3.4, [8.00, 7.73, 7.47, 7.07], 7.57),
]
for d, fhs, ft in data_511:
    f1,f2,f3,f4 = fhs
    rows.append(dict(d=d, h=0.3, a=0.0, n=1,
                     f1000=f1, f2000=f2, f3000=f3, f4000=f4, ft=ft))

# 表 5.1.2  单层，最优 d=1.6，a=0，n=1，随 h 变化
data_512 = [
    # h, [f1000,f2000,f3000,f4000], f_t
    (0.2, [8.80, 8.40, 8.00, 7.47], 8.17),
    (0.3, [10.53, 10.27, 9.87, 9.33], 9.99),
    (0.4, [10.93, 10.67, 10.40, 10.00], 10.50),
]
for h, fhs, ft in data_512:
    f1,f2,f3,f4 = fhs
    rows.append(dict(d=1.6, h=h, a=0.0, n=1,
                     f1000=f1, f2000=f2, f3000=f3, f4000=f4, ft=ft))

# 表 5.2  双层，h=0.3，n=1，a=d，随 d(=a) 变化
data_52 = [
    # d=a, [f1000,f2000,f3000,f4000], f_t
    (0.4, [8.13, 7.73, 7.33, 6.93], 7.53),
    (0.6, [9.60, 9.33, 8.93, 8.53], 9.10),
    (0.8, [11.73, 11.33, 10.93, 10.53], 11.13),
    (1.0, [11.07, 10.67, 10.27, 9.87], 10.47),
    (1.2, [10.27, 9.87, 9.47, 9.27], 9.67),
]
for da, fhs, ft in data_52:
    f1,f2,f3,f4 = fhs
    rows.append(dict(d=da, h=0.3, a=da, n=1,
                     f1000=f1, f2000=f2, f3000=f3, f4000=f4, ft=ft))

# 表 5.4  单层，d≈1.6，h=0.3，a=0，随 n 变化（按统一公式重新算 P）
data_54 = [
    # n, [f1000,f2000,f3000,f4000], f_t
    (1, [17.07, 16.27, 15.47, 14.67], 15.87),
    (3, [14.67, 14.00, 13.33, 12.67], 13.67),
    (5, [12.27, 11.60, 10.93, 10.27], 11.27),
    (7, [9.07, 8.40, 7.73, 7.07], 8.07),
    (9, [5.73, 5.07, 4.40, 3.73], 4.73),
]
for n, fhs, ft in data_54:
    f1,f2,f3,f4 = fhs
    rows.append(dict(d=1.6, h=0.3, a=0.0, n=n,
                     f1000=f1, f2000=f2, f3000=f3, f4000=f4, ft=ft))

df = pd.DataFrame(rows)

# -----------------------------
# 2) 计算 CV（四个频段的 f_h 的变异系数，百分数）
# -----------------------------
fh_cols = ["f1000", "f2000", "f3000", "f4000"]
df["fh_mean"] = df[fh_cols].mean(axis=1)
df["fh_std"]  = df[fh_cols].std(axis=1, ddof=1)
df["CV"]      = (df["fh_std"] / df["fh_mean"]) * 100.0  # %

# -----------------------------
# 3) 计算综合性能指数 P（按你给的权重公式）
# -----------------------------
df["P"] = (
    0.25 * df["ft"]
    + 0.20 * df["f1000"]
    + 0.20 * df["f2000"]
    + 0.18 * df["f3000"]
    + 0.17 * df["f4000"]
    + 0.23 * (1 - df["CV"] / 100.0)
)

print("\n数据预览：")
#print(df[["d","h","a","n","ft","f1000","f2000","f3000","f4000","CV","P"]].round(3).head())
print(df[["d","h","a","n","ft","f1000","f2000","f3000","f4000","CV","P"]].round(3))

# -----------------------------
# 4) 建模：P = f(d,h,a,n)  （多元二次多项式回归）
# -----------------------------
X = df[["d","h","a","n"]].values
y = df["P"].values

# poly = PolynomialFeatures(degree=1, include_bias=True)
# X_poly = poly.fit_transform(X)
# reg = LinearRegression().fit(X_poly, y)
reg = LinearRegression().fit(X, y)

# 训练 R^2
r2_train = reg.score(X_poly, y)
print("\n训练 R^2:", round(r2_train, 4))   # 注意：round() 用于 float

# 系数表
feature_names = poly.get_feature_names_out(["d","h","a","n"])
coef = np.append(reg.intercept_, reg.coef_[1:])  # 把截距放到 '1' 位置
coef_table = pd.DataFrame({"term": ["1"] + list(feature_names[1:]), "coef": coef})
print("\n模型系数:")
print(coef_table)

# 便捷预测函数
def predict_P(d, h, a, n):
    x = np.array([[d,h,a,n]])
    xp = poly.transform(x)
    return float(reg.predict(xp))

print("\n示例预测：d=1.6, h=0.3, a=0, n=1 -> P =",
      round(predict_P(1.6,0.3,0.0,1), 3))

# -----------------------------
# 5) 可视化
#   (1) d-a-P 三维曲面（h=0.3, n=1）
#   (2) h-P 曲线（d=1.6, a=0, n=1）
#   (3) n-P 曲线（d=1.6, h=0.3, a=0）
# -----------------------------
# (1) d-a-P surface
d_grid = np.linspace(0.4, 2.0, 50)
a_grid = np.linspace(0.0, 1.2, 50)
D, A = np.meshgrid(d_grid, a_grid)
H = np.full_like(D, 0.3)
N = np.full_like(D, 1.0)
Xg = np.column_stack([D.ravel(), H.ravel(), A.ravel(), N.ravel()])
Pg = reg.predict(poly.transform(Xg)).reshape(D.shape)

fig = plt.figure()
ax = fig.add_subplot(111, projection='3d')
ax.plot_surface(D, A, Pg, linewidth=0, antialiased=True)
ax.set_xlabel('d (cm)')
ax.set_ylabel('a (cm)')
ax.set_zlabel('Predicted P')
ax.set_title('P surface vs d & a (h=0.3, n=1)')
plt.tight_layout()
plt.savefig("plot_surface_d_a_P.png", dpi=200)
plt.show()

# (2) h-P 曲线（固定 d=1.6, a=0, n=1）
hs = np.linspace(0.2, 0.5, 60)
Ps_h = [predict_P(1.6, h, 0.0, 1) for h in hs]
plt.figure()
plt.plot(hs, Ps_h)
mask = (df["d"]==1.6) & (df["a"]==0) & (df["n"]==1)
plt.scatter(df.loc[mask, "h"], df.loc[mask, "P"])
plt.xlabel('h (cm)')
plt.ylabel('P')
plt.title('P vs h (d=1.6, a=0, n=1)')
plt.tight_layout()
plt.savefig("plot_curve_h_P.png", dpi=200)
plt.show()

# (3) n-P 曲线（固定 d=1.6, h=0.3, a=0）
ns = np.arange(1, 10, 1)
Ps_n = [predict_P(1.6, 0.3, 0.0, int(n)) for n in ns]
plt.figure()
plt.plot(ns, Ps_n)
mask = (df["d"]==1.6) & (df["h"]==0.3) & (df["a"]==0)
plt.scatter(df.loc[mask, "n"], df.loc[mask, "P"])
plt.xlabel('n (sources)')
plt.ylabel('P')
plt.title('P vs n (d=1.6, h=0.3, a=0)')
plt.tight_layout()
plt.savefig("plot_curve_n_P.png", dpi=200)
plt.show()

# 导出数据与系数（可选）
df.to_csv("dataset_with_P.csv", index=False)
coef_table.to_csv("poly2_linear_coef.csv", index=False)
print("\n已保存：plot_surface_d_a_P.png, plot_curve_h_P.png, plot_curve_n_P.png, dataset_with_P.csv, poly2_linear_coef.csv")


# =============================
# 6) 利用当前拟合模型，做一个网格搜索找近似最优解
#    限定在实验合理范围内：
#    d: 0.4 ~ 3.4
#    h: 0.2 ~ 0.4
#    a: 0.0 ~ 1.2
#    n: {1,3,5,7,9}
# =============================
def find_best_params():
    # 网格精度可以视情况调节（越密越准，但越慢）
    d_space = np.linspace(0.4, 3.4, 61)   # 步长约 0.05
    h_space = np.linspace(0.2, 0.4, 21)   # 步长约 0.01
    a_space = np.linspace(0.0, 1.2, 61)   # 步长约 0.02
    n_space = [1, 3, 5, 7, 9]             # 声源数量只取整数

    best_P = -np.inf
    best_params = None

    for d in d_space:
        for h in h_space:
            for a in a_space:
                for n in n_space:
                    P_val = predict_P(d, h, a, n)
                    if P_val > best_P:
                        best_P = P_val
                        best_params = (d, h, a, n)

    print("\n=== 近似最优解（基于网格搜索） ===")
    print(f"P_max ≈ {best_P:.3f}")
    print(f"d ≈ {best_params[0]:.3f} cm")
    print(f"h ≈ {best_params[1]:.3f} cm")
    print(f"a ≈ {best_params[2]:.3f} cm")
    print(f"n ≈ {int(best_params[3])} (sources)")
    return best_P, best_params

# 调用一次，打印结果
best_P, best_params = find_best_params()
