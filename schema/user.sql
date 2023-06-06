CREATE TABLE `anysock`.`user`
(
`id` INT(6) NOT NULL AUTO_INCREMENT ,
`username` VARCHAR NOT NULL ,
`password` VARCHAR NOT NULL ,
`email` TEXT NOT NULL ,
`balance` INT NOT NULL ,
PRIMARY KEY (`id`)
)
ENGINE = InnoDB;