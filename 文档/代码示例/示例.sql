/*--------Mysql数据库常用方法 开始--------*/
CASE 字段 WHEN 匹配 THEN 值 ELSE 默认值 END                                    // 当字段匹配时返回对应值，否则返回默认值
IF(条件, 值1, 值2)                                                             // 当条件为真，返回值1,否则返回值2
IFNULL(值1, 值2)                                                               // 当值1为null时，返回值2
NULLIF(值1, 值2)                                                               // 当值1等于值2，返回null，否则返回值1           //Postgresql通用
COALESCE(值,...)                                                               // 返回第一个非NULL参数                          //Postgresql通用

AVG([DISTINCT] 字段)                                                           // 平均值
ROUND(值, 2)                                                                   // 保留两位小数
CONCAT(字符串,...)                                                             // 拼接字符串                                    //Postgresql通用
CONCAT_WS(分隔符, 字符串,...)                                                  // 拼接字符串。可指定分隔符                      //Postgresql通用
GROUP_CONCAT([DISTINCT] 字段 [ORDER BY 排序字段 ASC/DESC] [SEPARATOR 分隔符])  // 拼接字符串。一般在GROUP BY语句中使用
REPLACE(字符串, ',', '')                                                       // 替换字符串。                                  //Postgresql通用
LENGTH(字符串)                                                                 // 字符串长度。中文根据编码不同，算多个字符
CHAR_LENGTH(字符串)                                                            // 字符串长度。中英文都算1个字符

UNIX_TIMESTAMP() 或 UNIX_TIMESTAMP('2006-01-02 15:04:05')                      // 时间戳。示例：1136185445
FROM_UNIXTIME(1136185445, '%Y-%m-%d %H:%i:%s')                                 // 时间戳转换成指定格式
DATE_FORMAT('2006-01-02 15:04:05', '%Y-%m-%d')                                 // 日期格式转换成指定格式
STR_TO_DATE('January 02 2016', '%M %d %Y')                                     // 根据指定格式将字符串转变成日期格式。示例：2006-01-02或2006-01-02 15:04:05
NOW() 或 UTC_TIMESTAMP()                                                       // 当前日期和时间。示例：2006-01-02 15:04:05
CURDATE() 或 UTC_DATE()                                                        // 当前日期。示例：2006-01-02
CURTIME() 或 UTC_TIME()                                                        // 当前时间。示例：15:04:05
YEAR('2006-01-02 15:04:05')                                                    // 年
MONTH('2006-01-02 15:04:05')                                                   // 月
DAY('2006-01-02 15:04:05')                                                     // 日
HOUR('2006-01-02 15:04:05')                                                    // 时
MINUTE('2006-01-02 15:04:05')                                                  // 分
SECOND('2006-01-02 15:04:05')                                                  // 秒
WEEK('2006-01-02 15:04:05')                                                    // 周。范围0~53
WEEKDAY('2006-01-02 15:04:05')                                                 // 周几。0星期一
LAST_DAY('2006-01-02 15:04:05')                                                // 返回当前日期月份的最后一天。示例：2006-01-31
DATE('2006-01-02 15:04:05')                                                    // 日期。示例：2006-01-02
DATE_SUB('2006-01-02 15:04:05', INTERVAL 7 type)                               // 该日期的多少type前。type：YEAR MONTH DAY HOUR MINUTE SECOND
DATE_ADD('2006-01-02 15:04:05', INTERVAL 7 type)                               // 该日期的多少type后。type：YEAR MONTH DAY HOUR MINUTE SECOND
DATEDIFF('2006-01-03', '2006-01-02')                                           // 相隔天数。示例：1
TIMEDIFF('16:05:06', '15:04:05')                                               // 相隔时间。示例：01:01:01

POINT(118.585519, 24.914168)                                                   // 字段类型为point时使用
ST_X(POINT(118.585519, 24.914168))                                             // 经度，Mysql5中使用X()
ST_Y(POINT(118.585519, 24.914168))                                             // 纬度，Mysql5中使用Y()
ST_DISTANCE_SPHERE(POINT(118.585519, 24.914168), POINT(118.585519, 24.914168)) // 计算经纬度距离

FIELD(字段, 值1, 值2,...)                                                      // 在ORDER BY后使用，字段按给定值进行排序。注意顺序问题，正序：方法内未列出的其它值, 值1, 值2,...；倒序：...,值2, 值1, 方法内未列出的其它值
/*--------Mysql数据库常用方法 结束--------*/
