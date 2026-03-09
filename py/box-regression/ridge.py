# -*- coding: utf-8 -*-
"""
多变量隔音建模 & 全局最优搜索
P = F(d, a, C)

- d: 内层隔音盒尺寸 (cm)
- a: 双层间距 (cm)
- C: 糖水浓度 (%)
- P: 综合隔音率 (%) —— 目标是最大化 P

用法：
1. 在 DATA 区域补全 / 修改你的实验数据。
2. 在 PARAM 区域修改 d, a, C 的搜索范围和步长。
3. 运行脚本，将打印出最优的 (d*, a*, C*) 和对应 P*。
"""

import numpy as np
from sklearn.preprocessing import PolynomialFeatures
from sklearn.linear_model import Ridge
from sklearn.pipeline import make_pipeline


# =========================
# 1. 实验数据区（请在这里补数据）
# =========================
# 每一行: [d, a, C, P]
# 你可以把自己的所有实验数据都填进来。
# 下面只先放我们对话里已经出现的几组，方便你看格式。
# 注意：单位统一：d(cm), a(cm), C(%), P(%)

data = np.array([
    # ---- 双层，空气（或未填充），外层 d_out = 21.5 ----
    # a = d_out - d_in, C = 0 (不加糖水)
    [15.0,  6.5,  0.0, 18.16568],
    [15.5,  6.0,  0.0, 16.40837],
    [16.1,  5.4,  0.0, 20.30286],
    [16.6,  4.9,  0.0, 17.62768],
    [17.2,  4.3,  0.0, 16.95930],
    [18.2,  3.3,  0.0,  9.75019],
    [19.0,  2.5,  0.0, 20.83073],
    [19.8,  1.7,  0.0, 15.13856],

    # ---- 双层 + 糖水填充，外层 d_out = 21.5, d_in = 19, a = 2.5 ----
    # C: 糖水浓度 (%)
    [19.0,  2.5,  1.0, 26.7949],
    [19.0,  2.5,  2.0, 32.1275],
    [19.0,  2.5,  4.0, 31.9052],
    [19.0,  2.5,  8.0, 28.8886],
    [19.0,  2.5, 16.0, 31.5802],

    # ---- （可选）把单层当作 a=0, C=0 的退化情况加入 ----
    # d: 15 ~ 22.1 的单层数据
    [15.0, 0.0, 0.0, 12.8750131],
    [15.5, 0.0, 0.0,  0.7765107],
    [16.1, 0.0, 0.0,  0.5235904],
    [16.6, 0.0, 0.0, -1.6066302],
    [17.2, 0.0, 0.0, -2.5766620],
    [18.2, 0.0, 0.0,  0.3169809],
    [19.0, 0.0, 0.0, 11.0882572],
    [19.8, 0.0, 0.0,  2.1120951],
    [20.7, 0.0, 0.0, 15.2052697],
    [21.5, 0.0, 0.0, 15.2498209],
    [22.1, 0.0, 0.0, 11.0418035],
])

# =========================
# 2. 模型参数区（可以自由调）
# =========================

# 多项式阶数：2 一般就够了，想更复杂可改成 3
POLY_DEGREE = 2

# 岭回归正则系数：越大表示惩罚越强，防止过拟合
RIDGE_ALPHA = 1.0

# 搜索范围 & 步长（这里就是你说的“通过修改 d,a,C 的取值范围”）
# 你可以自由调整：例如 d 在 18~21 搜，步长 0.1
# D_MIN, D_MAX, D_STEP = 18.0, 21.5, 0.1
# A_MIN, A_MAX, A_STEP = 1.5,  7.0, 0.1
# C_MIN, C_MAX, C_STEP = 0.0, 16.0, 0.5

D_MIN, D_MAX, D_STEP = 18.0, 30.0, 0.1
A_MIN, A_MAX, A_STEP = 1.5,  7.0, 0.1
C_MIN, C_MAX, C_STEP = 0.0, 30.0, 0.5


# 如果你希望强制 a = d_out - d_in 的几何关系，可在后面加上筛选逻辑
D_OUT = 21.5          # 固定外层尺寸（如需变化，也可以改成一个列表遍历）


# =========================
# 3. 训练模型
# =========================

X = data[:, :3]   # [d, a, C]
y = data[:, 3]    # P

model = make_pipeline(
    PolynomialFeatures(degree=POLY_DEGREE, include_bias=False),
    Ridge(alpha=RIDGE_ALPHA)
)
model.fit(X, y)

print("训练样本数:", len(X))


# =========================
# 4. 在给定范围内搜索最优解
# =========================

def search_best(d_range, a_range, c_range, constrain_geometry=False):
    """
    在给定范围内搜索 (d, a, C) 使得 P 最大。
    - d_range, a_range, c_range: np.ndarray
    - constrain_geometry: 若为 True，则强制 a ≈ D_OUT - d
    """
    best_P = -np.inf
    best_tuple = None

    for d in d_range:
        for a in a_range:
            # 可选几何约束：a 必须接近 D_OUT - d
            if constrain_geometry:
                if abs(a - (D_OUT - d)) > 0.2:  # 允许 ±0.2cm 的误差，你可以调整
                    continue

            for C in c_range:
                X_query = np.array([[d, a, C]])
                P_pred = model.predict(X_query)[0]
                if P_pred > best_P:
                    best_P = P_pred
                    best_tuple = (d, a, C)

    return best_tuple, best_P


def main():
    # 构造搜索网格
    d_range = np.arange(D_MIN, D_MAX + 1e-9, D_STEP)
    a_range = np.arange(A_MIN, A_MAX + 1e-9, A_STEP)
    c_range = np.arange(C_MIN, C_MAX + 1e-9, C_STEP)

    print("\n搜索范围：")
    print(f"  d ∈ [{D_MIN}, {D_MAX}]，步长 {D_STEP}")
    print(f"  a ∈ [{A_MIN}, {A_MAX}]，步长 {A_STEP}")
    print(f"  C ∈ [{C_MIN}, {C_MAX}]，步长 {C_STEP}")

    # ------ 1) 不加几何约束的全局搜索 ------
    best_params, best_P = search_best(d_range, a_range, c_range,
                                      constrain_geometry=False)
    d_star, a_star, C_star = best_params
    print("\n【全局搜索（不加几何约束）】")
    print(f"最优 d = {d_star:.3f} cm")
    print(f"最优 a = {a_star:.3f} cm")
    print(f"最优 C = {C_star:.3f} %")
    print(f"对应预测 P = {best_P:.3f} %")

    # ------ 2) 加几何约束：a ≈ D_OUT - d ------
    best_params_geo, best_P_geo = search_best(d_range, a_range, c_range,
                                              constrain_geometry=True)
    if best_params_geo is not None:
        d_star_g, a_star_g, C_star_g = best_params_geo
        print("\n【几何约束搜索（a ≈ d_out - d）】")
        print(f"外层 d_out = {D_OUT} cm")
        print(f"最优 d_in = {d_star_g:.3f} cm")
        print(f"最优 a = {a_star_g:.3f} cm (≈ d_out - d_in = {D_OUT - d_star_g:.3f})")
        print(f"最优 C = {C_star_g:.3f} %")
        print(f"对应预测 P = {best_P_geo:.3f} %")
    else:
        print("\n【几何约束搜索】未在给定范围内找到满足条件的解，请调整范围/误差限。")


if __name__ == "__main__":
    main()
