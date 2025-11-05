import numpy as np
import matplotlib.pyplot as plt
from mpl_toolkits.mplot3d import Axes3D  # noqa: F401
from sklearn.preprocessing import PolynomialFeatures
from sklearn.linear_model import LinearRegression
from matplotlib import font_manager

# =========================
# 0. 全局中文字体设置（按需修改路径）
# =========================
font_path = "/usr/share/fonts/opentype/noto/NotoSansCJK-Regular.ttc"
my_font = font_manager.FontProperties(fname=font_path)

plt.rcParams['axes.unicode_minus'] = False
plt.rcParams['font.family'] = my_font.get_name()


# =========================
# 1. 正方体俯视图
# =========================
def plot_cube_top_view():
    # 正方体顶点
    vertices = np.array([
        [0, 0, 0],  # 点 0
        [1, 0, 0],  # 点 1
        [1, 1, 0],  # 点 2
        [0, 1, 0],  # 点 3
        [0, 0, 1],  # 点 4
        [1, 0, 1],  # 点 5
        [1, 1, 1],  # 点 6
        [0, 1, 1],  # 点 7
    ])

    # 12 条边
    edges = [
        [0, 1], [1, 2], [2, 3], [3, 0],  # 底面
        [4, 5], [5, 6], [6, 7], [7, 4],  # 顶面
        [0, 4], [1, 5], [2, 6], [3, 7],  # 立柱
    ]

    fig = plt.figure(figsize=(6, 5))
    ax = fig.add_subplot(111, projection='3d')

    for edge in edges:
        ax.plot3D(*zip(*vertices[edge]), linewidth=1)

    # 俯视：从上往下看
    ax.view_init(elev=90, azim=45)
    ax.set_xlabel('X', fontproperties=my_font)
    ax.set_ylabel('Y', fontproperties=my_font)
    ax.set_zlabel('Z', fontproperties=my_font)
    ax.set_title("正方体的俯视图", fontproperties=my_font)

    plt.tight_layout()
    plt.savefig("cube_top_view.png", dpi=200, bbox_inches="tight")
    # plt.show()


# =========================
# 2. 拟合“内距/外距 -> 隔音率(%)”的 3D 曲面
# =========================
def fit_box_regression_surface():
    # X：列 0 = 内盒尺寸（音源到内层），列 1 = 外盒尺寸（音源到外层）
    X = np.array([
        [3, 5],
        [4, 7],
        [5, 9],
        [3, 9],
        [4, 9],
        [5, 9],
        [6, 9],
        [7, 9],
        [8, 9],
    ])

    # 原始测量的噪音值（dB）
    y_db = np.array([68.9, 70.5, 66.7, 72.2, 70.0, 68.5, 68.1, 69.7, 68.9])

    # 按公式换算成“隔音率(%)”
    # ratio = (75 - dB) / 75 * 100
    y = (75.0 - y_db) / 75.0 * 100.0

    # 多项式特征（二次项）
    poly = PolynomialFeatures(degree=2)
    X_poly = poly.fit_transform(X)

    model = LinearRegression()
    model.fit(X_poly, y)

    print("=== 拟合结果（目标 = 隔音率 %） ===")
    print("截距 intercept_:", model.intercept_)
    print("系数 coef_:", model.coef_)

    # 打印完整的多项式表达式（d = 内距, a = 外距）
    feature_names = poly.get_feature_names_out(['d', 'a'])
    terms = []
    for coef, name in zip(model.coef_, feature_names):
        terms.append(f"{coef:+.4f}*{name}")
    formula = f"y(%) = {model.intercept_:+.4f} " + " ".join(terms)
    formula = formula.replace('^2', '²').replace('*', '×')
    print("\n=== 拟合多项式（单位：隔音率 %） ===")
    print(formula)
    print("================================\n")

    # 预测函数：输入内距、外距，输出隔音率(%)
    def predict(inner_dist, outer_dist):
        arr = np.array([[inner_dist, outer_dist]], dtype=float)  # (1, 2)
        x_poly = poly.transform(arr)
        return model.predict(x_poly)

    # 构建网格
    inner_min, inner_max = 2, 9
    outer_min, outer_max = 2, 9
    x1_range = np.linspace(inner_min, inner_max, 100)  # 内
    x2_range = np.linspace(outer_min, outer_max, 100)  # 外
    X1, X2 = np.meshgrid(x1_range, x2_range)

    Z_flat = np.array([
        predict(x1, x2)[0] for x1, x2 in zip(X1.ravel(), X2.ravel())
    ])
    Z = Z_flat.reshape(X1.shape)  # Z 单位：隔音率 %

    return X, y, X1, X2, Z, predict


# =========================
# 3. 多视角 3D 曲面图（左右平视、中间 45°、俯视 45°）
# =========================
def plot_multi_view_surface(X1, X2, Z):
    fig = plt.figure(figsize=(12, 10))

    # 视角列表： (elev, azim)
    angles = [
        (0, -90),   # 从 X 轴正方向平视，突出“内层距离-隔音率”关系
        (0, 0),     # 从 Y 轴正方向平视，突出“外层距离-隔音率”关系
        (30, 45),   # 中间 45° 斜视
        (60, 45),   # 俯视 45°
    ]

    titles = [
        "平视（突出内层距离）",
        "平视（突出外层距离）",
        "斜视 45°",
        "俯视 45°",
    ]

    for i, (angle, title) in enumerate(zip(angles, titles), start=1):
        ax = fig.add_subplot(2, 2, i, projection='3d')
        ax.plot_surface(X1, X2, Z, edgecolor='none')
        ax.view_init(elev=angle[0], azim=angle[1])
        ax.set_xlabel('音源到内层盒的距离 (cm)', fontproperties=my_font)
        ax.set_ylabel('音源到外层盒的距离 (cm)', fontproperties=my_font)
        ax.set_zlabel('隔音率 (%)', fontproperties=my_font)
        ax.set_title(title, fontproperties=my_font)

    plt.tight_layout()
    plt.savefig("surface_multi_view.png", dpi=300, bbox_inches="tight")
    # plt.show()


# =========================
# 4. “俯视图”：X–Y 平面上的隔音率热力图
# =========================
def plot_top_heatmap(X1, X2, Z):
    fig, ax = plt.subplots(figsize=(6, 5))

    # 用等高线 / 伪彩色图表现俯视效果
    c = ax.contourf(X1, X2, Z, levels=20)
    fig.colorbar(c, ax=ax, label="预测隔音率 (%)")

    ax.set_xlabel("音源到内层盒的距离 (cm)", fontproperties=my_font)
    ax.set_ylabel("音源到外层盒的距离 (cm)", fontproperties=my_font)
    ax.set_title("俯视图：内层/外层距离平面上的隔音率", fontproperties=my_font)

    plt.tight_layout()
    plt.savefig("surface_top_view_heatmap.png", dpi=300, bbox_inches="tight")
    # plt.show()


# =========================
# 5. “内层距离–隔音率” & “外层距离–隔音率” 2D 关系图
# =========================
def plot_1d_relations(X, y_percent, predict):
    # y_percent 是隔音率(%)
    inner = X[:, 0]
    outer = X[:, 1]

    # 扫描内层距离，外层固定为最大值（比如 9 cm）
    inner_scan = np.linspace(inner.min(), inner.max(), 100)
    outer_fix = outer.max()
    y_inner_line = np.array([predict(xi, outer_fix) for xi in inner_scan]).ravel()

    # 扫描外层距离，内层取平均值
    outer_scan = np.linspace(outer.min(), outer.max(), 100)
    inner_fix = inner.mean()
    y_outer_line = np.array([predict(inner_fix, xo) for xo in outer_scan]).ravel()

    fig, axes = plt.subplots(1, 2, figsize=(12, 4))

    # 左图：内层距离
    ax1 = axes[0]
    ax1.scatter(inner, y_percent, label="实验数据", s=40)
    ax1.plot(inner_scan, y_inner_line, label=f"拟合曲线（外层 = {outer_fix:.1f} cm）")
    ax1.set_xlabel("音源到内层盒的距离 (cm)", fontproperties=my_font)
    ax1.set_ylabel("隔音率 (%)", fontproperties=my_font)
    ax1.set_title("内层距离与隔音率", fontproperties=my_font)
    ax1.grid(True)
    ax1.legend(prop=my_font)

    # 右图：外层距离
    ax2 = axes[1]
    ax2.scatter(outer, y_percent, label="实验数据", s=40)
    ax2.plot(outer_scan, y_outer_line, label=f"拟合曲线（内层 ≈ {inner_fix:.1f} cm）")
    ax2.set_xlabel("音源到外层盒的距离 (cm)", fontproperties=my_font)
    ax2.set_ylabel("隔音率 (%)", fontproperties=my_font)
    ax2.set_title("外层距离与隔音率", fontproperties=my_font)
    ax2.grid(True)
    ax2.legend(prop=my_font)

    plt.tight_layout()
    plt.savefig("inner_outer_1d_relations.png", dpi=300, bbox_inches="tight")
    # plt.show()


# =========================
# main：一次生成所有图
# =========================
if __name__ == "__main__":
    # 1. 正方体俯视图（作为“3D 空间坐标系”的直观示意）
    plot_cube_top_view()

    # 2. 拟合 3D 曲面（现在目标是隔音率 %）
    X, y_percent, X1, X2, Z, predict = fit_box_regression_surface()

    # 3. 多视角 3D 曲面图（左右平视 + 45° 斜视 + 45° 俯视）
    plot_multi_view_surface(X1, X2, Z)

    # 4. 俯视热力图（X–Y 平面上的隔音率分布）
    plot_top_heatmap(X1, X2, Z)

    # 5. 内/外层距离–隔音率 关系图（2D 投影）
    plot_1d_relations(X, y_percent, predict)

    print("所有图片已生成：")
    print("  cube_top_view.png")
    print("  surface_multi_view.png")
    print("  surface_top_view_heatmap.png")
    print("  inner_outer_1d_relations.png")
