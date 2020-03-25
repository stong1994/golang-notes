## 一、MyISAM与InnoDB区别

1. InnoDB支持事务，MyISAM不支持，对于InnoDB每一条SQL语言都默认封装成事务，自动提交，这样会影响速度，所以最好把多条SQL语言放在begin和commit之间，组成一个事务；

2. InnoDB支持外键，而MyISAM不支持。对一个包含外键的InnoDB表转为MYISAM会失败

3. InnoDB是**聚集索引**，使用**B+Tree**作为索引结构，数据文件是和（主键）索引绑在一起的（表数据文件本身就是按B+Tree组织的一个索引结构），必须要有主键，通过主键索引效率很高。但是辅助索引需要两次查询，先查询到主键，然后再通过主键查询到数据。因此，主键不应该过大，因为主键太大，其他索引也都会很大。

   MyISAM是**非聚集索引**，也是使用B+Tree作为索引结构，索引和数据文件是分离的，索引保存的是数据文件的指针。主键索引和辅助索引是独立的。

   也就是说：**InnoDB的B+树主键索引的叶子节点就是数据文件，辅助索引的叶子节点是主键的值；而MyISAM的B+树主键索引和辅助索引的叶子节点都是数据文件的地址指针。**

4. InnoDB不保存表的具体行数，执行select count(\*) from table时需要全表扫描。而MyISAM用一个变量保存了整个表的行数，执行上述语句时只需要读出该变量即可，速度很快（注意不能加有任何WHERE条件）；

5.  InnoDB支持表、行(默认)级锁，而MyISAM支持表级锁

6. InnoDB的行锁是实现在索引上的，而不是锁在物理行记录上。潜台词是，如果访问没有命中索引，也无法使用行锁，将要退化为表锁。

7. InnoDB表必须有主键（用户没有指定的话会自己找或生产一个主键），而Myisam可以没有

8. Innodb存储文件有frm、ibd，而Myisam是frm、MYD、MYI

   > Innodb：frm是表定义文件，ibd是数据文件
   >
   > Myisam：frm是表定义文件，myd是数据文件，myi是索引文件

#### 如何选择

   1. 是否要支持事务，如果要请选择innodb，如果不需要可以考虑MyISAM；
   2. 如果表中绝大多数都只是读查询，可以考虑MyISAM，如果既有读也有写，请使用InnoDB。
   3. 系统奔溃后，MyISAM恢复起来更困难，能否接受；
   4. MySQL5.5版本开始Innodb已经成为Mysql的默认引擎(之前是MyISAM)，说明其优势是有目共睹的，如果你不知道用什么，那就用InnoDB，至少不会差。 

#### InnoDB为什么推荐使用自增ID作为主键？

`自增ID可以保证每次插入时B+索引是从右边扩展的，可以避免B+树和频繁合并和分裂（对比使用UUID）。如果使用字符串主键和随机主键，会使得数据随机插入，效率比较差。`

#### innodb引擎的4大特性

`插入缓冲（insert buffer),二次写(double write),自适应哈希索引(ahi),预读(read ahead)`

#### 参考文章

1. https://blog.csdn.net/qq_35642036/article/details/82820178
## 二. 如何查看一个sql有没有命中索引
在sql语句前增加`explain`即可，

主要字段|解释
---|---
table|显示这一行的数据是关于哪张表的
type|显示连接使用了何种类型。从最好到最差的连接类型为const、eq_reg、ref、range、index和ALL
possible_keys|显示可能应用在这张表中的索引。如果为空，没有可能的索引。
key|实际使用的索引。如果为NULL，则没有使用索引
key_len|使用的索引的长度。在不损失精确性的情况下，长度越短越好
ref|显示索引的哪一列被使用了，如果可能的话，是一个常数
rows|MYSQL认为必须检查的用来返回请求数据的行数
Extra|关于MYSQL如何解析查询的额外信息。如果是Using temporary和Using filesort，意思MYSQL根本不能使用索引，结果是检索会很慢

extra列返回的描述的意义： 
1. Distinct:一旦MYSQL找到了与行相联合匹配的行，就不再搜索了
2. Not exists: MYSQL优化了LEFT JOIN，一旦它找到了匹配LEFT JOIN标准的行，就不再搜索了
3. Range checked for each Record（index map:#）:没有找到理想的索引，因此对于从前面表中来的每一个行组合，MYSQL检查使用哪个索引，并用它来从表中返回行。这是使用索引的最慢的连接之一
4. Using filesort: 看到这个的时候，查询就需要优化了。MYSQL需要进行额外的步骤来发现如何对返回的行排序。它根据连接类型以及存储排序键值和匹配条件的全部行的行指针来排序全部行
5. Using index: 列数据是从仅仅使用了索引中的信息而没有读取实际的行动的表返回的，这发生在对表的全部的请求列都是同一个索引的部分的时候 
6. Using temporary 看到这个的时候，查询需要优化了。这里，MYSQL需要创建一个临时表来存储结果，这通常发生在对不同的列集进行ORDER BY上，而不是GROUP BY上

[explain详解](https://www.jianshu.com/p/ea3fc71fdc45)

3. 数据库事务的四大特征https://blog.csdn.net/xlgen157387/article/details/79450295
    > ACID: 原子性、一致性、隔离性、持久性
4. char和varchar的区别
    > 
5. mysql为什么使用B+树
    > https://blog.csdn.net/xlgen157387/article/details/79450295
6. 自增ID对高并发有什么影响