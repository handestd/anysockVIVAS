-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Jun 08, 2023 at 12:51 PM
-- Server version: 10.4.28-MariaDB
-- PHP Version: 8.2.4

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `anysock2`
--

-- --------------------------------------------------------

--
-- Table structure for table `user`
--

CREATE TABLE `user` (
  `id` int(6) NOT NULL,
  `username` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `email` text NOT NULL,
  `balance` int(6) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

--
-- Dumping data for table `user`
--

INSERT INTO `user` (`id`, `username`, `password`, `email`, `balance`) VALUES
(1, 'tiendat0', '123456', 'handestd@gmail.com', 5),
(2, 'tiendat', '$2a$14$Y7UbVywVEmGqvyXLsk/o/uLz40sOiyfrl8kDCuD2aqrz3V71CCj8G', 'handestd@gmail.com', 0),
(4, 'tiendat1', '$2a$14$ZhE08LggACmWtnvOiy.bku1miBbtecAhIE6YI0Gwtyf.LG8d6Mzh6', 'handestd@gmail.com', 0),
(5, 'tiendat2', '$2a$14$LPgRnUCHZXUJdGfvtPtunOdB2JKpfluyiEJaC.fDgiVtHovDNyps2', 'handestd@gmail.com', 0),
(6, 'tiendat3', '$2a$14$8.Mqi4hTV5mASScVA6gRt.PnvJIzMEqHEy0hsFfTDeTuKpjHWFXyK', 'handestd@gmail.com', 0),
(7, 'tiendat4', '$2a$14$5p3UGLiVUudJ0jTrqQ4Cx.coYRx3wPiMUJ7xSpsSJslyfioVaQjt.', 'handestd@gmail.com', 0);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `user`
--
ALTER TABLE `user`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `user`
--
ALTER TABLE `user`
  MODIFY `id` int(6) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=8;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
