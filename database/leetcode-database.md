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
