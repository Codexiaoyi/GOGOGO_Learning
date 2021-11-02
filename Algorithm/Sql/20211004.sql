-- 182
select Email from 
(
    select Email,count(Email) as num
    from Person
    group by Email
) as t
where num > 1
;
-- 