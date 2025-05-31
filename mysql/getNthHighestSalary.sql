/* 177. Nth Highest Salary*/



CREATE FUNCTION getNthHighestSalary(N INT) RETURNS INT
BEGIN
    SET N = N-1;
  RETURN (
    SELECT DISTINCT salary
    FROM Employee
    ORDER BY salary DESC
    LIMIT N,1
  );
END

/*runtime 486ms Beats 28.65%