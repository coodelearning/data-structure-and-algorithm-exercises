# 175. 组合两个表 https://leetcode-cn.com/problems/combine-two-tables/
SELECT  `FirstName`
       ,`LastName`
       ,`City`
       ,`State`
FROM Person
LEFT JOIN Address
ON Address.PersonId = Person.PersonId;

# 182. 查找重复的电子邮箱 https://leetcode-cn.com/problems/duplicate-emails/
# Function 1:

# Step1:获得临时表
SELECT  Email
       ,COUNT(Email)
FROM Person
GROUP BY  Email;
# Step2:
SELECT  Email
FROM
(
	SELECT  Email
	       ,COUNT(Email) AS num
	FROM Person
	GROUP BY  Email
) AS statistic
WHERE num >1;

# Function 2:

SELECT  Email
FROM Person
GROUP BY  Email
HAVING COUNT(Email)>1;
