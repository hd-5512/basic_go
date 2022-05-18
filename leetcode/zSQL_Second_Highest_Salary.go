package leetcode

import "fmt"

/**
编写一个 SQL 查询，获取 Employee 表中第二高的薪水（Salary） 。

+----+--------+
| Id | Salary |
+----+--------+
| 1  | 100    |
| 2  | 200    |
| 3  | 300    |
+----+--------+
例如上述 Employee 表，SQL查询应该返回 200 作为第二高的薪水。如果不存在第二高的薪水，那么查询应返回 null。

+---------------------+
| SecondHighestSalary |
+---------------------+
| 200                 |
+---------------------+

 */

func zSQL_Second_Highest_Salary() {

	sql := "SELECT " +
				"CASE COUNT(*) " +
				"WHEN 1 THEN NULL " +
				"WHEN 1 THEN NULL " +
				"ELSE MIN(s.Salary) END as SecondHighestSalary " +
			"FROM ( " +
				"SELECT Salary FROM Employee GROUP BY Salary ORDER BY Salary DESC LIMIT 2 " +
			") as s"

	fmt.Println(sql)
}
