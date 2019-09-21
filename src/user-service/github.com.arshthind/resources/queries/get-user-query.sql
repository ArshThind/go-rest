select user_id   as userID,
       user_name as userName,
       email,
       phone_num as phoneNum,
       user_Type as userType
from users
where status = 'A'