-- phpMyAdmin SQL Dump
-- version 5.1.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: May 10, 2022 at 09:21 AM
-- Server version: 10.4.21-MariaDB
-- PHP Version: 8.0.11

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `db_user_product`
--

-- --------------------------------------------------------

--
-- Table structure for table `products`
--

CREATE TABLE `products` (
  `id` int(11) NOT NULL,
  `name` varchar(255) DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL,
  `status` tinyint(1) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `products`
--

INSERT INTO `products` (`id`, `name`, `description`, `status`) VALUES
(1, 'Product Test', 'lorem ipsum lorem ipsum lorem ipsum lorem ipsum', 0),
(2, 'Product Test', 'lorem ipsum lorem ipsum lorem ipsum lorem ipsum', 0),
(3, 'Product Test', 'lorem ipsum lorem ipsum lorem ipsum lorem ipsum', 0),
(4, 'Product Test', 'lorem ipsum lorem ipsum lorem ipsum lorem ipsum', 1),
(5, 'Product Test', 'lorem ipsum lorem ipsum lorem ipsum lorem ipsum', 1),
(6, 'Product Test', 'lorem ipsum lorem ipsum lorem ipsum lorem ipsum', 1),
(7, 'Product Test', 'lorem ipsum lorem ipsum lorem ipsum lorem ipsum', 1),
(8, 'Product Test', 'lorem ipsum lorem ipsum lorem ipsum lorem ipsum', 1),
(9, 'Product Test', 'lorem ipsum lorem ipsum lorem ipsum lorem ipsum', 1),
(10, 'Product Test', 'lorem ipsum lorem ipsum lorem ipsum lorem ipsum', 1),
(11, 'Product Test', 'lorem ipsum lorem ipsum lorem ipsum lorem ipsum', 1),
(12, 'Product Test', 'lorem ipsum lorem ipsum lorem ipsum lorem ipsum', 1),
(13, 'Product Test', 'lorem ipsum lorem ipsum lorem ipsum lorem ipsum', 1),
(14, 'Product Test', 'lorem ipsum lorem ipsum lorem ipsum lorem ipsum', 1),
(15, 'Product Test', 'lorem ipsum lorem ipsum lorem ipsum lorem ipsum', 1),
(16, 'Product Test', 'lorem ipsum lorem ipsum lorem ipsum lorem ipsum', 1),
(17, 'Product Test', 'lorem ipsum lorem ipsum lorem ipsum lorem ipsum', 1),
(18, 'Product Test', 'lorem ipsum lorem ipsum lorem ipsum lorem ipsum', 1),
(19, 'Product Test', 'lorem ipsum lorem ipsum lorem ipsum lorem ipsum', 1),
(20, 'Product Test', 'lorem ipsum lorem ipsum lorem ipsum lorem ipsum', 1),
(21, 'Product Test', 'lorem ipsum lorem ipsum lorem ipsum lorem ipsum', 1),
(22, 'Product Test', 'lorem ipsum lorem ipsum lorem ipsum lorem ipsum', 1),
(23, 'Product Test', 'lorem ipsum lorem ipsum lorem ipsum lorem ipsum', 1),
(24, 'Product Test', 'lorem ipsum lorem ipsum lorem ipsum lorem ipsum', 1),
(25, 'Product Test', 'lorem ipsum lorem ipsum lorem ipsum lorem ipsum', 1),
(26, 'Product Test', 'lorem ipsum lorem ipsum lorem ipsum lorem ipsum', 1),
(27, 'Product Test', 'lorem ipsum lorem ipsum lorem ipsum lorem ipsum', 1),
(28, 'Product Test', 'lorem ipsum lorem ipsum lorem ipsum lorem ipsum', 1),
(29, 'Product Test', 'lorem ipsum lorem ipsum lorem ipsum lorem ipsum', 1),
(30, 'Product Test', 'lorem ipsum lorem ipsum lorem ipsum lorem ipsum', 1);

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` int(11) NOT NULL,
  `username` varchar(255) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `phone` varchar(255) DEFAULT NULL,
  `role` varchar(255) DEFAULT NULL,
  `status` tinyint(1) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `username`, `email`, `phone`, `role`, `status`) VALUES
(2, 'Admin Test', 'admin@gmail.com', '0812344366', 'admin', 1),
(3, 'Admin Test 123', 'admin@gmail.com', '0812344366', 'admin', 1),
(4, 'Admin Test', 'admin@gmail.com', '0812344366', 'admin', 1),
(5, 'Admin Test', 'admin@gmail.com', '0812344366', 'admin', 1),
(6, 'Admin Test', 'admin@gmail.com', '0812344366', 'admin', 1),
(7, 'Admin Test', 'admin@gmail.com', '0812344366', 'admin', 1),
(8, 'Admin Test', 'admin@gmail.com', '0812344366', 'admin', 1),
(9, 'Admin Test', 'admin@gmail.com', '0812344366', 'admin', 1),
(10, 'Admin Test', 'admin@gmail.com', '0812344366', 'admin', 1),
(11, 'Admin Test', 'admin@gmail.com', '0812344366', 'admin', 1),
(12, 'Admin Test', 'admin@gmail.com', '0812344366', 'admin', 1),
(13, 'Admin Test', 'admin@gmail.com', '0812344366', 'admin', 1),
(14, 'Admin Test', 'admin@gmail.com', '0812344366', 'admin', 1),
(15, 'Admin Test', 'admin@gmail.com', '0812344366', 'admin', 1),
(16, 'Admin Test', 'admin@gmail.com', '0812344366', 'admin', 1),
(17, 'Admin Test', 'admin@gmail.com', '0812344366', 'admin', 1),
(18, 'Admin Test', 'admin@gmail.com', '0812344366', 'admin', 1),
(19, 'Admin Test', 'admin@gmail.com', '0812344366', 'admin', 1),
(20, 'Admin Test', 'admin@gmail.com', '0812344366', 'admin', 0);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `products`
--
ALTER TABLE `products`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `products`
--
ALTER TABLE `products`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=31;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=21;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
