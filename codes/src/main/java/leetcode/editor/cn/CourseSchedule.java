/**
 * 你这个学期必须选修 numCourses 门课程，记为 0 到 numCourses - 1 。
 * <p>
 * 在选修某些课程之前需要一些先修课程。 先修课程按数组 prerequisites 给出，其中 prerequisites[i] = [ai, bi] ，表示如
 * 果要学习课程 ai 则 必须 先学习课程 bi 。
 * <p>
 * <p>
 * 例如，先修课程对 [0, 1] 表示：想要学习课程 0 ，你需要先完成课程 1 。
 * <p>
 * <p>
 * 请你判断是否可能完成所有课程的学习？如果可以，返回 true ；否则，返回 false 。
 * <p>
 * <p>
 * <p>
 * 示例 1：
 * <p>
 * <p>
 * 输入：numCourses = 2, prerequisites = [[1,0]]
 * 输出：true
 * 解释：总共有 2 门课程。学习课程 1 之前，你需要完成课程 0 。这是可能的。
 * <p>
 * 示例 2：
 * <p>
 * <p>
 * 输入：numCourses = 2, prerequisites = [[1,0],[0,1]]
 * 输出：false
 * 解释：总共有 2 门课程。学习课程 1 之前，你需要先完成​课程 0 ；并且学习课程 0 之前，你还应先完成课程 1 。这是不可能的。
 * <p>
 * <p>
 * <p>
 * 提示：
 * <p>
 * <p>
 * 1 <= numCourses <= 10⁵
 * 0 <= prerequisites.length <= 5000
 * prerequisites[i].length == 2
 * 0 <= ai, bi < numCourses
 * prerequisites[i] 中的所有课程对 互不相同
 * <p>
 * <p>
 * Related Topics 深度优先搜索 广度优先搜索 图 拓扑排序 👍 1595 👎 0
 */

package leetcode.editor.cn;

public class CourseSchedule {

    public static void main(String[] args) {
        Solution solution = new CourseSchedule().new Solution();
        System.out.println(solution.canFinish(2, new int[][]{{1, 0}}));
    }

    //leetcode submit region begin(Prohibit modification and deletion)
    class Solution {
        public boolean canFinish(int numCourses, int[][] prerequisites) {
            return false;
        }
    }
    //leetcode submit region end(Prohibit modification and deletion)

}

