package leetcode


func zSQL_DuplicateEmails() string {
	return ""
}
/*
	SELECT a.Email FROM  (
    	SELECT Email,count(*) as cnt FROM Person group by Email
 	) as a WHERE a.cnt > 1
*/


/*
	SELECT Email FROM Person GROUP BY Email HAVING count(*) > 1
 */
