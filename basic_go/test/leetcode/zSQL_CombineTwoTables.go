package leetcode

/*
表1: Person

+-------------+---------+
| 列名         | 类型     |
+-------------+---------+
| PersonId    | int     |
| FirstName   | varchar |
| LastName    | varchar |
+-------------+---------+
PersonId 是上表主键
表2: Address

+-------------+---------+
| 列名         | 类型    |
+-------------+---------+
| AddressId   | int     |
| PersonId    | int     |
| City        | varchar |
| State       | varchar |
+-------------+---------+
AddressId 是上表主键
 

编写一个 SQL 查询，满足条件：无论 person 是否有地址信息，都需要基于上述两表提供 person 的以下信息：

FirstName, LastName, City, State

 */

func zSQL_CombineTwoTables() {
	/*
		SELECT
			a.FirstName,a.LastName,b.City,b.State
		FROM Person as a
			LEFT JOIN Address as b on a.PersonId = b.PersonId
	*/
}


/**********************************************************************************************************/

//考点 join 中 最怕的就两项  一个是 关联顺序错误  导致业务错误  结果溢出 或 拦截

//第二个就是 效率
