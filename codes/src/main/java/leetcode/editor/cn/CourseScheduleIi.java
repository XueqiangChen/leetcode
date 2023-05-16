/**
 * 现在你总共有 numCourses 门课需要选，记为 0 到 numCourses - 1。给你一个数组 prerequisites ，其中
 * prerequisites[i] = [ai, bi] ，表示在选修课程 ai 前 必须 先选修 bi 。
 * <p>
 * <p>
 * 例如，想要学习课程 0 ，你需要先完成课程 1 ，我们用一个匹配来表示：[0,1] 。
 * <p>
 * <p>
 * 返回你为了学完所有课程所安排的学习顺序。可能会有多个正确的顺序，你只要返回 任意一种 就可以了。如果不可能完成所有课程，返回 一个空数组 。
 * <p>
 * <p>
 * <p>
 * 示例 1：
 * <p>
 * <p>
 * 输入：numCourses = 2, prerequisites = [[1,0]]
 * 输出：[0,1]
 * 解释：总共有 2 门课程。要学习课程 1，你需要先完成课程 0。因此，正确的课程顺序为 [0,1] 。
 * <p>
 * <p>
 * 示例 2：
 * <p>
 * <p>
 * 输入：numCourses = 4, prerequisites = [[1,0],[2,0],[3,1],[3,2]]
 * 输出：[0,2,1,3]
 * 解释：总共有 4 门课程。要学习课程 3，你应该先完成课程 1 和课程 2。并且课程 1 和课程 2 都应该排在课程 0 之后。
 * 因此，一个正确的课程顺序是 [0,1,2,3] 。另一个正确的排序是 [0,2,1,3] 。
 * <p>
 * 示例 3：
 * <p>
 * <p>
 * 输入：numCourses = 1, prerequisites = []
 * 输出：[0]
 * <p>
 * <p>
 * <p>
 * 提示：
 * <p>
 * <p>
 * 1 <= numCourses <= 2000
 * 0 <= prerequisites.length <= numCourses * (numCourses - 1)
 * prerequisites[i].length == 2
 * 0 <= ai, bi < numCourses
 * ai != bi
 * 所有[ai, bi] 互不相同
 * <p>
 * <p>
 * Related Topics 深度优先搜索 广度优先搜索 图 拓扑排序 👍 772 👎 0
 */

package leetcode.editor.cn;

import java.util.ArrayList;
import java.util.LinkedList;
import java.util.List;
import java.util.Queue;

public class CourseScheduleIi {
    public static void main(String[] args) {
        Solution solution = new CourseScheduleIi().new Solution();
        System.out.println(solution.findOrder(2, new int[][]{{1,0}})[1]);
    }

    //leetcode submit region begin(Prohibit modification and deletion)
    class Solution {

        // 存储有向图，key: 要修的课程， value: 先修的课程列表
        List<List<Integer>> edges;
        // 存储每个节点的入度
        int[] indeg;
        // 存储答案
        int[] result;
        // 答案下标
        int index;

        public int[] findOrder(int numCourses, int[][] prerequisites) {
            edges = new ArrayList<>();
            for (int i = 0; i < numCourses; i++) {
                edges.add(new ArrayList<>());
            }
            indeg = new int[numCourses];
            result = new int[numCourses];
            index = 0;
            // [1, 0] 学习课程 1 之前，要完成课程 0， 拓扑图即是: 0->1, 1 的入度+1，
            for (int[] info : prerequisites) {
                edges.get(info[1]).add(info[0]);
                // 用下标来表示不同节点的入度，对于字符串的数组来说，是行不通的，字符串数组要用map来表示节点与度之间的关系
                ++indeg[info[0]];
            }

            Queue<Integer> queue = new LinkedList<>();
            // 将所有入度为 0 的节点放入队列中
            for (int i = 0; i < numCourses; i++) {
                if (indeg[i] == 0) {
                    queue.offer(i);
                }
            }

            while (!queue.isEmpty()) {
                // 从队首取出一个节点
                int u = queue.poll();
                // 放入答案中
                result[index++] = u;
                for (int v : edges.get(u)) {
                    --indeg[v];
                    // 如果相邻节点 v 的入度为 0， 就可以选 v 对应的课程了
                    if (indeg[v] == 0) {
                        queue.offer(v);
                    }
                }
            }

            // 有环存在
            if (index != numCourses) {
                return new int[0];
            }

            return result;
        }
    }
    //leetcode submit region end(Prohibit modification and deletion)

    // 通过深度优先搜索的方式来遍历有向无环图
    class Solution2 {
        // 存储有向图
        List<List<Integer>> edges;

        // 标记每个节点的状态：0=未搜索，1=搜索中，2=已完成
        int[] visited;
        // 用数组来模拟栈，下标 n-1 为栈底，0 为栈顶
        int[] result;
        // 判断有向图中是否有环
        boolean valid = true;
        // 栈下标
        int index;

        public int[] findOrder(int numCourses, int[][] prerequisites) {
            edges = new ArrayList<List<Integer>>();
            for (int i = 0; i < numCourses; ++i) {
                edges.add(new ArrayList<>());
            }
            visited = new int[numCourses];
            result = new int[numCourses];
            index = numCourses - 1;
            for (int[] info : prerequisites) {
                edges.get(info[1]).add(info[0]);
            }
            // 每次挑选一个「未搜索」的节点，开始进行深度优先搜索
            for (int i = 0; i < numCourses && valid; ++i) {
                if (visited[i] == 0) {
                    dfs(i);
                }
            }
            if (!valid) {
                return new int[0];
            }
            // 如果没有环，那么就有拓扑排序
            return result;
        }

        public void dfs(int u) {
            // 将节点标记为「搜索中」
            visited[u] = 1;
            // 搜索其相邻节点
            // 只要发现有环，立刻停止搜索
            for (int v : edges.get(u)) {
                // 如果「未搜索」那么搜索相邻节点
                if (visited[v] == 0) {
                    dfs(v);
                    if (!valid) {
                        return;
                    }
                }
                // 如果「搜索中」说明找到了环
                else if (visited[v] == 1) {
                    valid = false;
                    return;
                }
            }
            // 将节点标记为「已完成」
            visited[u] = 2;
            // 将节点入栈
            result[index--] = u;
        }

    }

}

