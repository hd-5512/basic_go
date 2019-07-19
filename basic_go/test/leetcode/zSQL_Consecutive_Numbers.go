package leetcode


/**
编写一个 SQL 查询，查找所有至少连续出现三次的数字。

+----+-----+
| Id | Num |
+----+-----+
| 1  |  1  |
| 2  |  1  |
| 3  |  1  |
| 4  |  2  |
| 5  |  1  |
| 6  |  2  |
| 7  |  2  |
+----+-----+
例如，给定上面的 Logs 表， 1 是唯一连续出现至少三次的数字。

+-----------------+
| ConsecutiveNums |
+-----------------+
| 1               |
+-----------------+

 */

func zSQL_Consecutive_Numbers(){
	/**
	SELECT DISTINCT
	    l1.Num AS ConsecutiveNums
	FROM
	    Logs l1,
	    Logs l2,
	    Logs l3
	WHERE
	    l1.Id = l2.Id - 1
	    AND l2.Id = l3.Id - 1
	    AND l1.Num = l2.Num
	    AND l2.Num = l3.Num
	*/


/*
	select
		distinct a.Num as ConsecutiveNums
	from
		(select Num,
		case
		when  @recorde = Num  then @count := @count+1
			when  @recorde <> @recorde:= Num then @count :=1 end as n
			from Logs,
				(
					select @count := 0,@recorde := (select Num from Logs limit 0,1)
				)r
		)a
	where a.n >= 3
*/
}

