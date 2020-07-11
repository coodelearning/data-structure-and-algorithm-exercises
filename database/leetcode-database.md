# 数据库练习

## 简单

### 175. [组合两个表](https://leetcode-cn.com/problems/combine-two-tables/)

```mssql
SELECT  `FirstName`
       ,`LastName`
       ,`City`
       ,`State`
FROM Person
LEFT JOIN Address
ON Address.PersonId = Person.PersonId;
```


### 182. [查找重复的电子邮箱](https://leetcode-cn.com/problems/duplicate-emails/)

#### Function 1:

```mssql
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
```


#### Function 2:

* 向 `GROUP BY` 添加条件的一种更常用的方法是使用 `HAVING` 子句，该子句更为简单高效。

  所以我们可以将上面的解决方案重写为：

```mysql
SELECT  Email
FROM Person
GROUP BY  Email
HAVING COUNT(Email)>1;
```

### [595. 大的国家](https://leetcode-cn.com/problems/big-countries/)

```mysql
SELECT  name
       ,population
       ,area
FROM World
WHERE area > 3000000 OR population > 25000000
```

### [596. 超过5名学生的课](https://leetcode-cn.com/problems/classes-more-than-5-students/)

```mysql
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
```