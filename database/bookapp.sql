-- phpMyAdmin SQL Dump
-- version 5.1.1
-- https://www.phpmyadmin.net/
--
-- Máy chủ: 127.0.0.1
-- Thời gian đã tạo: Th8 01, 2022 lúc 05:02 AM
-- Phiên bản máy phục vụ: 10.4.22-MariaDB
-- Phiên bản PHP: 8.1.1

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Cơ sở dữ liệu: `bookapp`
--

-- --------------------------------------------------------

--
-- Cấu trúc bảng cho bảng `author`
--

CREATE TABLE `author` (
  `author_id` int(11) NOT NULL,
  `name` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `phone` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Đang đổ dữ liệu cho bảng `author`
--

INSERT INTO `author` (`author_id`, `name`, `email`, `phone`, `password`) VALUES
(11, 'a', 'a@gmail.com', '1234567890', '$2a$10$EGWSPltnBTNPCoRp0XyzA.P4RTZov8ocx1wAs0ElSxB6p9evolapy'),
(12, 'oke', 'zzz@gmail.com', '9876543210', '$2a$04$mmKUP66mbxRTHzFFyPqTO.X/7zdvORcXDBIVS/cZOwPFTPR5NHgtK'),
(13, 'ppppp', 'r@gmail.com', '', '$2a$04$mmKUP66mbxRTHzFFyPqTO.X/7zdvORcXDBIVS/cZOwPFTPR5NHgtK'),
(15, 'quang', 'b@gmail.com', '', '$2a$04$mmKUP66mbxRTHzFFyPqTO.X/7zdvORcXDBIVS/cZOwPFTPR5NHgtK'),
(16, 'quang', 'quang.nguyen@sotatek.com', '', '$2a$04$mmKUP66mbxRTHzFFyPqTO.X/7zdvORcXDBIVS/cZOwPFTPR5NHgtK'),
(18, 'quang', 'quangnguyen@sotatek.com', '', '$2a$10$0plDZb5r0TkBh8mLzcYZtexP3yModCJsAxaJFqXsHx17pNHVAjzXW'),
(19, 'quangnn2k1', 'qc@sotatek.com', '01234567898', '$2a$10$ii2inQw70IelwMllz8MzlucCFAnQO6vlbibh.Xz6g4i8aviWTfjmG'),
(20, 'quang', 'q@sotatek.com', '0123456789', '$2a$10$mmR2yFja6DLOrW8YV4FZgePF4pDcZJ6JQ7/McmNbNR1oHHqHrJXw2'),
(21, 'quang', 'qq@sotatek.com', '01234567899', '$2a$10$1.25sgwQgnn17IMkAni4cOIYDH9bH9QME6WYTNse.hw0pScfmvN2u'),
(22, 'quangnn', 'quang2k1@sotatek.com', '01234567891', '$2a$10$qg3LJx8hkjdTdgyjd7CODOyRxypBj./J6mekoM6/1l.q3FLD.vQcq'),
(23, 'quangnn', 'quang2k@sotatek.com', '01234567892', '$2a$10$EGWSPltnBTNPCoRp0XyzA.P4RTZov8ocx1wAs0ElSxB6p9evolapy');

-- --------------------------------------------------------

--
-- Cấu trúc bảng cho bảng `book`
--

CREATE TABLE `book` (
  `book_id` int(11) NOT NULL,
  `name` text NOT NULL,
  `description` text NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Đang đổ dữ liệu cho bảng `book`
--

INSERT INTO `book` (`book_id`, `name`, `description`) VALUES
(1, 'abc', '1'),
(5, 'bca', '10'),
(10, 'mbappe', '10'),
(19, 'fds', '2'),
(22, 'reee', '19a'),
(23, '11111', ''),
(26, 'ccc', ''),
(28, 'quang', 'eeeeee'),
(29, 'hoan', 'hhhhh'),
(30, 'oke', 'zzz'),
(31, 'zk', '<3'),
(32, 'here', ''),
(33, 'jjjjjjjjjjjjjj', '');

-- --------------------------------------------------------

--
-- Cấu trúc bảng cho bảng `book_author`
--

CREATE TABLE `book_author` (
  `book_id` int(11) NOT NULL,
  `author_id` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Đang đổ dữ liệu cho bảng `book_author`
--

INSERT INTO `book_author` (`book_id`, `author_id`) VALUES
(1, 11),
(1, 12),
(1, 13),
(5, 11),
(5, 12),
(5, 13);

-- --------------------------------------------------------

--
-- Cấu trúc bảng cho bảng `book_category`
--

CREATE TABLE `book_category` (
  `book_id` int(11) NOT NULL,
  `category_id` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Đang đổ dữ liệu cho bảng `book_category`
--

INSERT INTO `book_category` (`book_id`, `category_id`) VALUES
(1, 1),
(1, 2),
(1, 3),
(5, 1),
(5, 2),
(5, 3),
(5, 4),
(1, 4),
(22, 1);

-- --------------------------------------------------------

--
-- Cấu trúc bảng cho bảng `category`
--

CREATE TABLE `category` (
  `category_id` int(11) NOT NULL,
  `name` varchar(255) NOT NULL,
  `description` text NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Đang đổ dữ liệu cho bảng `category`
--

INSERT INTO `category` (`category_id`, `name`, `description`) VALUES
(1, 'ppppp', ''),
(2, '1322', ''),
(3, 'zzzzzzz', ''),
(4, 'llllll', '');

--
-- Chỉ mục cho các bảng đã đổ
--

--
-- Chỉ mục cho bảng `author`
--
ALTER TABLE `author`
  ADD PRIMARY KEY (`author_id`);

--
-- Chỉ mục cho bảng `book`
--
ALTER TABLE `book`
  ADD PRIMARY KEY (`book_id`);

--
-- Chỉ mục cho bảng `book_author`
--
ALTER TABLE `book_author`
  ADD KEY `book_id` (`book_id`),
  ADD KEY `author_id` (`author_id`);

--
-- Chỉ mục cho bảng `book_category`
--
ALTER TABLE `book_category`
  ADD KEY `book_id` (`book_id`),
  ADD KEY `category_id` (`category_id`);

--
-- Chỉ mục cho bảng `category`
--
ALTER TABLE `category`
  ADD PRIMARY KEY (`category_id`);

--
-- AUTO_INCREMENT cho các bảng đã đổ
--

--
-- AUTO_INCREMENT cho bảng `author`
--
ALTER TABLE `author`
  MODIFY `author_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=24;

--
-- AUTO_INCREMENT cho bảng `book`
--
ALTER TABLE `book`
  MODIFY `book_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=34;

--
-- AUTO_INCREMENT cho bảng `category`
--
ALTER TABLE `category`
  MODIFY `category_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- Các ràng buộc cho các bảng đã đổ
--

--
-- Các ràng buộc cho bảng `book_author`
--
ALTER TABLE `book_author`
  ADD CONSTRAINT `book_author_ibfk_1` FOREIGN KEY (`book_id`) REFERENCES `book` (`book_id`),
  ADD CONSTRAINT `book_author_ibfk_2` FOREIGN KEY (`author_id`) REFERENCES `author` (`author_id`);

--
-- Các ràng buộc cho bảng `book_category`
--
ALTER TABLE `book_category`
  ADD CONSTRAINT `book_category_ibfk_1` FOREIGN KEY (`book_id`) REFERENCES `book` (`book_id`),
  ADD CONSTRAINT `book_category_ibfk_2` FOREIGN KEY (`category_id`) REFERENCES `category` (`category_id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
