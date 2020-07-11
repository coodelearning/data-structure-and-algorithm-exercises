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

# 595. 大的国家 https://leetcode-cn.com/problems/big-countries/
SELECT  name
       ,population
       ,area
FROM World
WHERE area > 3000000 OR population > 25000000

# 596
# Function 1:
SELECT  class
FROM
(
	SELECT  class
	       ,COUNT(student) AS num
	FROM courses
	GROUP BY  class
) AS statistic
WHERE num > 4;
# Function 2:
SELECT  class
FROM courses
GROUP BY  class
HAVING COUNT(student) > 4;