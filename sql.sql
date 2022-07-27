/*
SQLyog Ultimate v11.24 (32 bit)
MySQL - 8.0.26 : Database - pethome
*********************************************************************
*/

/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
CREATE DATABASE /*!32312 IF NOT EXISTS*/`pethome` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;

USE `pethome`;

/*Table structure for table `cart_items` */

DROP TABLE IF EXISTS `cart_items`;

CREATE TABLE `cart_items` (
                              `ID` int NOT NULL AUTO_INCREMENT,
                              `COUNT` int NOT NULL,
                              `amount` double(11,2) NOT NULL,
  `service_id` int NOT NULL,
  `cart_id` varchar(100) NOT NULL,
  PRIMARY KEY (`ID`),
  KEY `car_id` (`service_id`),
  KEY `cart_id` (`cart_id`),
  CONSTRAINT `cart_items_ibfk_1` FOREIGN KEY (`service_id`) REFERENCES `services` (`ID`),
  CONSTRAINT `cart_items_ibfk_2` FOREIGN KEY (`cart_id`) REFERENCES `carts` (`ID`)
) ENGINE=InnoDB AUTO_INCREMENT=49 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Data for the table `cart_items` */

/*Table structure for table `carts` */

DROP TABLE IF EXISTS `carts`;

CREATE TABLE `carts` (
                         `ID` varchar(100) NOT NULL,
                         `total_count` int NOT NULL,
                         `total_amount` double(11,2) NOT NULL,
  `user_id` int NOT NULL,
  PRIMARY KEY (`ID`),
  KEY `user_id` (`user_id`),
  CONSTRAINT `carts_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Data for the table `carts` */

/*Table structure for table `order_items` */

DROP TABLE IF EXISTS `order_items`;

CREATE TABLE `order_items` (
                               `id` int NOT NULL AUTO_INCREMENT,
                               `count` int NOT NULL,
                               `amount` double(11,2) NOT NULL,
  `name` varchar(100) NOT NULL,
  `price` double(11,2) NOT NULL,
  `order_id` varchar(100) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `order_id` (`order_id`),
  CONSTRAINT `order_items_ibfk_1` FOREIGN KEY (`order_id`) REFERENCES `orders` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=34 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Data for the table `order_items` */

insert  into `order_items`(`id`,`count`,`amount`,`name`,`price`,`order_id`) values (21,1,1300.00,'猫猫寄养',1300.00,'95af5a25-3679-41ba-62ff-6CD471C483F1'),(22,1,300.00,'仓鼠寄养',300.00,'95af5a25-3679-41ba-62ff-6CD471C483F1'),(23,1,110.00,'进口猫粮',110.00,'95af5a25-3679-41ba-62ff-6CD471C483F1'),(24,1,1300.00,'猫猫寄养',1300.00,'6325253f-ec73-4dd7-69e2-8BF921119C16'),(25,1,300.00,'仓鼠寄养',300.00,'6325253f-ec73-4dd7-69e2-8BF921119C16'),(26,2,220.00,'进口猫粮',110.00,'6325253f-ec73-4dd7-69e2-8BF921119C16'),(27,2,160.00,'宠物洗澡',80.00,'6325253f-ec73-4dd7-69e2-8BF921119C16'),(28,3,3900.00,'猫猫寄养',1300.00,'0bf50598-7592-4e66-4a5b-DF2C7FC48445'),(29,1,110.00,'进口猫粮',110.00,'0bf50598-7592-4e66-4a5b-DF2C7FC48445'),(30,2,400.00,'国产狗粮',200.00,'0bf50598-7592-4e66-4a5b-DF2C7FC48445'),(31,1,300.00,'仓鼠寄养',300.00,'81855ad8-681d-4d86-51e9-1E00167939CB'),(32,6,180.00,'鱼肉干',30.00,'81855ad8-681d-4d86-51e9-1E00167939CB'),(33,3,240.00,'宠物洗澡',80.00,'81855ad8-681d-4d86-51e9-1E00167939CB');

/*Table structure for table `orders` */

DROP TABLE IF EXISTS `orders`;

CREATE TABLE `orders` (
                          `id` varchar(100) NOT NULL,
                          `create_time` datetime NOT NULL,
                          `total_count` int NOT NULL,
                          `total_amount` double(11,2) NOT NULL,
  `state` int NOT NULL,
  `user_id` int DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`),
  CONSTRAINT `orders_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Data for the table `orders` */

insert  into `orders`(`id`,`create_time`,`total_count`,`total_amount`,`state`,`user_id`) values ('0bf50598-7592-4e66-4a5b-DF2C7FC48445','2022-05-15 21:48:53',6,4410.00,0,2),('6325253f-ec73-4dd7-69e2-8BF921119C16','2022-05-15 21:47:46',6,1980.00,0,2),('81855ad8-681d-4d86-51e9-1E00167939CB','2022-05-15 22:40:04',7,480.00,0,1),('95af5a25-3679-41ba-62ff-6CD471C483F1','2022-05-15 21:35:09',3,1710.00,0,1);

/*Table structure for table `services` */

DROP TABLE IF EXISTS `services`;

CREATE TABLE `services` (
                            `ID` int NOT NULL AUTO_INCREMENT,
                            `Name` varchar(100) NOT NULL,
                            `Price` double(11,2) NOT NULL,
  `Num` int DEFAULT NULL,
  `ImgPath` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci,
  PRIMARY KEY (`ID`),
  UNIQUE KEY `Name` (`Name`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Data for the table `services` */

insert  into `services`(`ID`,`Name`,`Price`,`Num`,`ImgPath`) values (1,'猫猫寄养',1300.00,10,'https://gimg2.baidu.com/image_search/src=http%3A%2F%2Fc-ssl.duitang.com%2Fuploads%2Fitem%2F201905%2F25%2F20190525232957_cABaa.jpeg&refer=http%3A%2F%2Fc-ssl.duitang.com&app=2002&size=f9999,10000&q=a80&n=0&g=0n&fmt=auto?sec=1654700773&t=17d59e1b817de6db65cb73d02f9cc281'),(3,'仓鼠寄养',260.00,15,'https://img1.baidu.com/it/u=727761736,357832122&fm=253&fmt=auto&app=120&f=JPEG?w=800&h=800'),(4,'宠物洗澡',80.00,17,'https://img2.baidu.com/it/u=3777673847,4130984660&fm=253&fmt=auto&app=138&f=JPEG?w=400&h=266'),(5,'进口猫粮',123.00,15,'https://img0.baidu.com/it/u=3380399648,3708344995&fm=253&fmt=auto&app=138&f=JPEG?w=665&h=500'),(7,'逗猫棒',15.00,9,'https://img0.baidu.com/it/u=816595926,4116453711&fm=253&fmt=auto&app=138&f=JPEG?w=500&h=500'),(8,'宠物驱虫',331.00,15,'http://t15.baidu.com/it/u=3534739966,1706162033&fm=224&app=112&f=JPEG?w=500&h=500'),(9,'狗狗磨牙棒',15.00,10,'https://gimg2.baidu.com/image_search/src=http%3A%2F%2Fimg.alicdn.com%2Fbao%2Fuploaded%2Fi1%2F735355352%2FO1CN01p34P9k1pPGPsQNLL3_%21%210-item_pic.jpg&refer=http%3A%2F%2Fimg.alicdn.com&app=2002&size=f9999,10000&q=a80&n=0&g=0n&fmt=auto?sec=1655215544&t=d4514c27cde6fa1203421b1af1f687eb'),(11,'鱼肉干',30.00,60,'https://img0.baidu.com/it/u=1295697401,1586404818&fm=253&fmt=auto&app=138&f=PNG?w=500&h=500'),(12,'狗狗寄养',500.00,5,'https://img2.baidu.com/it/u=2305806164,1673211166&fm=253&fmt=auto&app=120&f=JPEG?w=1140&h=749');

/*Table structure for table `sessions` */

DROP TABLE IF EXISTS `sessions`;

CREATE TABLE `sessions` (
                            `SessionID` varchar(100) NOT NULL,
                            `Name` varchar(100) NOT NULL,
                            `UserID` int NOT NULL,
                            PRIMARY KEY (`SessionID`),
                            KEY `UserID` (`UserID`),
                            CONSTRAINT `sessions_ibfk_1` FOREIGN KEY (`UserID`) REFERENCES `users` (`ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Data for the table `sessions` */

insert  into `sessions`(`SessionID`,`Name`,`UserID`) values ('6694d2c4-22ac-4208-6007-2939487F6999','root',2),('9566c74d-1003-4c4d-7bbb-0407D1E2C649','root',1),('eb9d18a4-4784-445d-47f3-C67CF22746E9','root',1);

/*Table structure for table `users` */

DROP TABLE IF EXISTS `users`;

CREATE TABLE `users` (
                         `ID` int NOT NULL AUTO_INCREMENT,
                         `Name` varchar(100) NOT NULL,
                         `Age` int NOT NULL,
                         `Phone` varchar(100) NOT NULL,
                         `Username` varchar(100) NOT NULL,
                         `Password` varchar(100) NOT NULL,
                         PRIMARY KEY (`ID`),
                         UNIQUE KEY `Name` (`Name`),
                         UNIQUE KEY `Phone` (`Phone`),
                         UNIQUE KEY `Username` (`Username`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Data for the table `users` */

insert  into `users`(`ID`,`Name`,`Age`,`Phone`,`Username`,`Password`) values (1,'root',66,'666','root','111111'),(2,'test',2,'33','张三','111111'),(3,'测试',13,'1331313','李四','123456'),(4,'王五',18,'120120120','test123','123456'),(5,'懒羊羊',66,'666666','test66','1111111');

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
