package leetcode

/**
编写一个 SQL 查询，获取 Employee 表中第 n 高的薪水（Salary）。

+----+--------+
| Id | Salary |
+----+--------+
| 1  | 100    |
| 2  | 200    |
| 3  | 300    |
+----+--------+
例如上述 Employee 表，n = 2 时，应返回第二高的薪水 200。如果不存在第 n 高的薪水，那么查询应返回 null。

+------------------------+
| getNthHighestSalary(2) |
+------------------------+
| 200                    |
+------------------------+

 */

func zSQL_Nth_Highest_Salary(){
	/**
	CREATE FUNCTION getNthHighestSalary(N INT) RETURNS INT
	BEGIN
	  RETURN (
	      # Write your MySQL query statement below.
	      	SELECT
	            CASE 1
	        WHEN COUNT(*) BETWEEN 0 and (N-1) THEN NULL
	        ELSE MIN(s.Salary) END as `getNthHighestSalary(N)`
	        FROM (
					SELECT Salary FROM Employee GROUP BY Salary ORDER BY Salary DESC LIMIT N
			) as s
	  );
	END
	 */
}
