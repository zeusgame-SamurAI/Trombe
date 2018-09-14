package algorithm

func astar() {

}

/* Python A* algorithm
import heapq
import itertools


def astar(init_pos, goal):
    # 探索した座標を格納する経路リスト
    passed_list = [init_pos]
    # 初期スコア
    init_score = distance(passed_list) + heuristic(init_pos)
    # 探索済み座標と、その座標に辿り着いた経路のスコアを格納
    checked = {init_pos: init_score}
    # 経路リストとそのスコアを格納する探索ヒープ
    searching_heap = []
    # 探索ヒープに経路リストを格納
    heapq.heappush(searching_heap, (init_score, passed_list))
    # 探索不可能になるまで
    while len(searching_heap) > 0:
        # 現在までに探索した経路の中から、スコアが最小になる
        # ときのスコアと経路リストを取得
        score, passed_list = heapq.heappop(searching_heap)
        # 最後に探索した座標
        last_passed_pos = passed_list[-1]
        # 最後に探索した座標が目的地なら探索ヒープを返す
        if last_passed_pos == goal:
            return passed_list
        # 最後に探索した座標の八方を探索
        for pos in nexts(last_passed_pos):
            # 経路リストに探索中の座標を追加した一時リストを作成
            new_passed_list = passed_list + [pos]
            # 一時リストのスコアを計算
            pos_score = distance(new_passed_list) + heuristic(pos)
            # 探索中の座標が、他の経路で探索済みかどうかチェック
            # 探索済みの場合、前回のスコアと今回のスコアを比較
            # 今回のスコアのほうが大きい場合、次の方角の座標の探索へ
            if pos in checked and checked[pos] <= pos_score:
                continue
            # 今回のスコアのほうが小さい場合、チェック済みリストに格納
            # 探索ヒープにスコアと経路リストを格納
            checked[pos] = pos_score
            heapq.heappush(searching_heap, (pos_score, new_passed_list))

    return []

if __name__ == "__main__":
    dungeon = [
        'OOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOO',
        'OS  O     O     O         O          O',
        'O   O  O  O  O            O    OOOO GO',
        'O      O     O  O   OOOO  O    O  OOOO',
        'OOOOOOOOOOOOOOOOOO  O     O    O     O',
        'O                O  O     O          O',
        'O        OOO     O  O     OOOOOOOOO  O',
        'O  OO    O    OOOO  O     O      OO  O',
        'O   O    O          O     O  O   O   O',
        'O   OOO  O          O        O   O   O',
        'O        O          O        O       O',
        'OOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOO',
        ]

    def find_ch(ch):
        for i, l in enumerate(dungeon):

            for j, c in enumerate(l):

                if c == ch:
                    return (i, j)
    # スタート
    init = find_ch("S")
    # ゴール
    goal = find_ch("G")

    def nexts(pos):
        ''' 探索中の座標から八方の座標を計算するジェネレーター'''
        wall = "O"

        for a, b in itertools.product([' + 1', ' - 1', ''], repeat=2):

            if a or b:

                if dungeon[eval('pos[0]' + a)][eval('pos[1]' + b)] != wall:
                    yield (eval('pos[0]' + a), eval('pos[1]' + b))

    def heuristic(pos):
        ''' 探索中の座標からゴールまでの最短距離のスコア '''
        return ((pos[0] - goal[0]) ** 2 + (pos[1] - goal[1]) ** 2) ** 0.5

    def distance(path):
        ''' スタートから探索中の座標までの距離のスコア '''
        return len(path)

    def render_path(path):
        ''' 結果の出力 '''
        buf = [[ch for ch in l] for l in dungeon]

        for pos in path[1:-1]:
            buf[pos[0]][pos[1]] = "*"

        buf[path[0][0]][path[0][1]] = "s"
        buf[path[-1][0]][path[-1][1]] = "g"
        return ["".join(l) for l in buf]

    path = astar(init, goal)

    if len(path) > 0:
        print("\n".join(render_path(path)))

    else:
		print('failed')
*/
