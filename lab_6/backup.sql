-- PostgreSQL database dump
-- Wed, 30 Apr 2025 11:22:16 +03

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;


--
-- Table: dinosaur
--

CREATE TABLE dinosaur (id integer NOT NULL, name character varying NOT NULL, species character varying NOT NULL, gender character varying NOT NULL, max_speed real, weight real, height real, enclosure_id integer);

INSERT INTO dinosaur (id, name, species, gender, max_speed, weight, height, enclosure_id) VALUES ('21', 'Дейзи', 'Птеродактиль', 'Женский', '65', '35', '0.7', '15');
INSERT INTO dinosaur (id, name, species, gender, max_speed, weight, height, enclosure_id) VALUES ('24', 'Оливер', 'Цератозавр', 'Мужской', '30', '1200', '2', '12');
INSERT INTO dinosaur (id, name, species, gender, max_speed, weight, height, enclosure_id) VALUES ('1', 'Джек', 'Тиранозавр', 'Мужской', '55.3', '7800', '4.1', '26');
INSERT INTO dinosaur (id, name, species, gender, max_speed, weight, height, enclosure_id) VALUES ('2', 'Ларри', 'Диплодок', 'Женский', '60', '6700', '5.6', '9');
INSERT INTO dinosaur (id, name, species, gender, max_speed, weight, height, enclosure_id) VALUES ('3', 'Рекс', 'Тиранозавр', 'Мужской', '50.5', '8000', '4.2', '13');
INSERT INTO dinosaur (id, name, species, gender, max_speed, weight, height, enclosure_id) VALUES ('4', 'Стелла', 'Трицератопс', 'Женский', '30', '6000', '3', '18');
INSERT INTO dinosaur (id, name, species, gender, max_speed, weight, height, enclosure_id) VALUES ('5', 'Бруно', 'Велоцираптор', 'Мужской', '65', '150', '1.8', '2');
INSERT INTO dinosaur (id, name, species, gender, max_speed, weight, height, enclosure_id) VALUES ('6', 'Луна', 'Стегозавр', 'Женский', '25', '5000', '4.5', '9');
INSERT INTO dinosaur (id, name, species, gender, max_speed, weight, height, enclosure_id) VALUES ('7', 'Спарки', 'Птеродактиль', 'Мужской', '80', '50', '1.2', '25');
INSERT INTO dinosaur (id, name, species, gender, max_speed, weight, height, enclosure_id) VALUES ('8', 'Тиша', 'Брахиозавр', 'Женский', '20', '30000', '12', '18');
INSERT INTO dinosaur (id, name, species, gender, max_speed, weight, height, enclosure_id) VALUES ('9', 'Гром', 'Анкилозавр', 'Мужской', '15', '6000', '2.5', '17');
INSERT INTO dinosaur (id, name, species, gender, max_speed, weight, height, enclosure_id) VALUES ('11', 'Рокки', 'Аллозавр', 'Мужской', '55', '2000', '3.5', '4');
INSERT INTO dinosaur (id, name, species, gender, max_speed, weight, height, enclosure_id) VALUES ('12', 'Мия', 'Галлимим', 'Женский', '70', '450', '2', '13');
INSERT INTO dinosaur (id, name, species, gender, max_speed, weight, height, enclosure_id) VALUES ('14', 'Эльза', 'Дилофозавр', 'Женский', '45', '400', '2.2', '5');
INSERT INTO dinosaur (id, name, species, gender, max_speed, weight, height, enclosure_id) VALUES ('15', 'Зевс', 'Спинозавр', 'Мужской', '35', '7000', '5', '25');
INSERT INTO dinosaur (id, name, species, gender, max_speed, weight, height, enclosure_id) VALUES ('17', 'Блейд', 'Дейноних', 'Мужской', '75', '100', '1.5', '15');
INSERT INTO dinosaur (id, name, species, gender, max_speed, weight, height, enclosure_id) VALUES ('18', 'Лили', 'Овираптор', 'Женский', '50', '80', '1', '27');
INSERT INTO dinosaur (id, name, species, gender, max_speed, weight, height, enclosure_id) VALUES ('19', 'Макс', 'Мозазавр', 'Мужской', '40', '15000', '10', '9');
INSERT INTO dinosaur (id, name, species, gender, max_speed, weight, height, enclosure_id) VALUES ('20', 'Нора', 'Археоптерикс', 'Женский', '60', '1', '0.5', '16');
INSERT INTO dinosaur (id, name, species, gender, max_speed, weight, height, enclosure_id) VALUES ('22', 'Гизмо', 'Протоцератопс', 'Мужской', '25', '400', '1.8', '8');
INSERT INTO dinosaur (id, name, species, gender, max_speed, weight, height, enclosure_id) VALUES ('23', 'Руби', 'Цератозавр', 'Женский', '45', '1000', '2.5', '27');
INSERT INTO dinosaur (id, name, species, gender, max_speed, weight, height, enclosure_id) VALUES ('25', 'Астра', 'Эдмонтозавр', 'Женский', '35', '3000', '3.5', '6');
INSERT INTO dinosaur (id, name, species, gender, max_speed, weight, height, enclosure_id) VALUES ('27', 'Чарли', 'Троодон', 'Мужской', '70', '50', '1', '7');
INSERT INTO dinosaur (id, name, species, gender, max_speed, weight, height, enclosure_id) VALUES ('10', 'Зоя', 'Паразауролоф', 'Мужской', '40', '2500', '3.8', '20');
INSERT INTO dinosaur (id, name, species, gender, max_speed, weight, height, enclosure_id) VALUES ('26', 'Бакс', 'Диплодок', 'Мужской', '20', '2500', '3', '25');
INSERT INTO dinosaur (id, name, species, gender, max_speed, weight, height, enclosure_id) VALUES ('13', 'Тор', 'Диплодок', 'Мужской', '60', '1500', '3', '17');
INSERT INTO dinosaur (id, name, species, gender, max_speed, weight, height, enclosure_id) VALUES ('16', 'Флора', 'Спинозавр', 'Женский', '30', '5700', '3.2', '13');
CREATE UNIQUE INDEX dinosaur_pkey ON public.dinosaur USING btree (id);


--
-- Table: employee
--

CREATE TABLE employee (id integer NOT NULL, name character varying NOT NULL, gender character varying NOT NULL, age integer, position character varying NOT NULL, experience integer);

INSERT INTO employee (id, name, gender, age, position, experience) VALUES ('1', 'Иван Петров', 'Мужской', '35', 'Смотритель', '10');
INSERT INTO employee (id, name, gender, age, position, experience) VALUES ('2', 'Мария Иванова', 'Женский', '28', 'Ветеринар', '5');
INSERT INTO employee (id, name, gender, age, position, experience) VALUES ('4', 'Елена Кузнецова', 'Женский', '32', 'Биолог', '8');
INSERT INTO employee (id, name, gender, age, position, experience) VALUES ('5', 'Дмитрий Волков', 'Мужской', '29', 'Охранник', '6');
INSERT INTO employee (id, name, gender, age, position, experience) VALUES ('6', 'Ольга Белова', 'Женский', '27', 'Администратор', '4');
INSERT INTO employee (id, name, gender, age, position, experience) VALUES ('10', 'Татьяна Лебедева', 'Женский', '26', 'Лаборант', '3');
INSERT INTO employee (id, name, gender, age, position, experience) VALUES ('11', 'Николай Соколов', 'Мужской', '38', 'Смотритель', '12');
INSERT INTO employee (id, name, gender, age, position, experience) VALUES ('12', 'Юлия Воробьева', 'Женский', '31', 'Ветеринар', '6');
INSERT INTO employee (id, name, gender, age, position, experience) VALUES ('13', 'Андрей Павлов', 'Мужской', '34', 'Инженер', '10');
INSERT INTO employee (id, name, gender, age, position, experience) VALUES ('15', 'Артем Михайлов', 'Мужской', '42', 'Охранник', '14');
INSERT INTO employee (id, name, gender, age, position, experience) VALUES ('16', 'Екатерина Федорова', 'Женский', '36', 'Администратор', '11');
INSERT INTO employee (id, name, gender, age, position, experience) VALUES ('17', 'Виктор Григорьев', 'Мужской', '39', 'Директор', '18');
INSERT INTO employee (id, name, gender, age, position, experience) VALUES ('19', 'Игорь Захаров', 'Мужской', '37', 'Техник', '13');
INSERT INTO employee (id, name, gender, age, position, experience) VALUES ('20', 'Светлана Николаева', 'Женский', '33', 'Лаборант', '8');
INSERT INTO employee (id, name, gender, age, position, experience) VALUES ('21', 'Роман Борисов', 'Мужской', '31', 'Смотритель', '7');
INSERT INTO employee (id, name, gender, age, position, experience) VALUES ('22', 'Марина Алексеева', 'Женский', '27', 'Ветеринар', '3');
INSERT INTO employee (id, name, gender, age, position, experience) VALUES ('23', 'Вадим Тимофеев', 'Мужской', '44', 'Инженер', '16');
INSERT INTO employee (id, name, gender, age, position, experience) VALUES ('25', 'Георгий Макаров', 'Мужской', '35', 'Охранник', '10');
INSERT INTO employee (id, name, gender, age, position, experience) VALUES ('7', 'Сергей Новиков', 'Мужской', '45', 'Охранник', '20');
INSERT INTO employee (id, name, gender, age, position, experience) VALUES ('8', 'Анна Морозова', 'Женский', '30', 'Техник', '7');
INSERT INTO employee (id, name, gender, age, position, experience) VALUES ('9', 'Павел Козлов', 'Мужской', '33', 'Биолог', '9');
INSERT INTO employee (id, name, gender, age, position, experience) VALUES ('14', 'Наталья Семенова', 'Женский', '29', 'Смотритель', '5');
INSERT INTO employee (id, name, gender, age, position, experience) VALUES ('18', 'Людмила Егорова', 'Женский', '28', 'Биолог', '4');
INSERT INTO employee (id, name, gender, age, position, experience) VALUES ('27', 'Станислав Орлов', 'Мужской', '41', 'Охранник', '17');
INSERT INTO employee (id, name, gender, age, position, experience) VALUES ('26', 'Мария Дмитриева', 'Женский', '29', 'Администратор', '5');
INSERT INTO employee (id, name, gender, age, position, experience) VALUES ('3', 'Андрей Смирнов', 'Мужской', '40', 'Инженер', '15');
INSERT INTO employee (id, name, gender, age, position, experience) VALUES ('24', 'Алексей Степанов', 'Мужской', '30', 'Смотритель', '6');
CREATE UNIQUE INDEX employee_pkey ON public.employee USING btree (id);


--
-- Table: employee_exhibit
--

CREATE TABLE employee_exhibit (employee_id integer NOT NULL, exhibit_id integer NOT NULL);

INSERT INTO employee_exhibit (employee_id, exhibit_id) VALUES ('1', '2');
INSERT INTO employee_exhibit (employee_id, exhibit_id) VALUES ('2', '8');
INSERT INTO employee_exhibit (employee_id, exhibit_id) VALUES ('3', '14');
INSERT INTO employee_exhibit (employee_id, exhibit_id) VALUES ('4', '19');
INSERT INTO employee_exhibit (employee_id, exhibit_id) VALUES ('5', '6');
INSERT INTO employee_exhibit (employee_id, exhibit_id) VALUES ('6', '11');
INSERT INTO employee_exhibit (employee_id, exhibit_id) VALUES ('7', '3');
INSERT INTO employee_exhibit (employee_id, exhibit_id) VALUES ('8', '15');
INSERT INTO employee_exhibit (employee_id, exhibit_id) VALUES ('9', '1');
INSERT INTO employee_exhibit (employee_id, exhibit_id) VALUES ('10', '22');
INSERT INTO employee_exhibit (employee_id, exhibit_id) VALUES ('11', '27');
INSERT INTO employee_exhibit (employee_id, exhibit_id) VALUES ('12', '4');
INSERT INTO employee_exhibit (employee_id, exhibit_id) VALUES ('13', '10');
INSERT INTO employee_exhibit (employee_id, exhibit_id) VALUES ('14', '25');
INSERT INTO employee_exhibit (employee_id, exhibit_id) VALUES ('15', '9');
INSERT INTO employee_exhibit (employee_id, exhibit_id) VALUES ('16', '17');
INSERT INTO employee_exhibit (employee_id, exhibit_id) VALUES ('17', '5');
INSERT INTO employee_exhibit (employee_id, exhibit_id) VALUES ('18', '12');
INSERT INTO employee_exhibit (employee_id, exhibit_id) VALUES ('19', '20');
INSERT INTO employee_exhibit (employee_id, exhibit_id) VALUES ('20', '24');
INSERT INTO employee_exhibit (employee_id, exhibit_id) VALUES ('21', '16');
INSERT INTO employee_exhibit (employee_id, exhibit_id) VALUES ('22', '21');
INSERT INTO employee_exhibit (employee_id, exhibit_id) VALUES ('23', '7');
INSERT INTO employee_exhibit (employee_id, exhibit_id) VALUES ('24', '18');
INSERT INTO employee_exhibit (employee_id, exhibit_id) VALUES ('25', '13');
INSERT INTO employee_exhibit (employee_id, exhibit_id) VALUES ('26', '26');
INSERT INTO employee_exhibit (employee_id, exhibit_id) VALUES ('27', '23');


--
-- Table: enclosure
--

CREATE TABLE enclosure (id integer NOT NULL, square real NOT NULL, location character varying NOT NULL, height real, capacity integer);

INSERT INTO enclosure (id, square, location, height, capacity) VALUES ('2', '300', 'Южная зона', '8', '3');
INSERT INTO enclosure (id, square, location, height, capacity) VALUES ('3', '450', 'Восточная зона', '9.5', '4');
INSERT INTO enclosure (id, square, location, height, capacity) VALUES ('4', '600', 'Западная зона', '12', '6');
INSERT INTO enclosure (id, square, location, height, capacity) VALUES ('5', '350', 'Центральная зона', '7', '2');
INSERT INTO enclosure (id, square, location, height, capacity) VALUES ('6', '400', 'Северо-Восточная зона', '8.5', '4');
INSERT INTO enclosure (id, square, location, height, capacity) VALUES ('7', '550', 'Юго-Западная зона', '11', '5');
INSERT INTO enclosure (id, square, location, height, capacity) VALUES ('8', '480', 'Северо-Западная зона', '9', '3');
INSERT INTO enclosure (id, square, location, height, capacity) VALUES ('9', '700', 'Юго-Восточная зона', '13', '7');
INSERT INTO enclosure (id, square, location, height, capacity) VALUES ('21', '850', 'Юго-Северная зона', '17', '13');
INSERT INTO enclosure (id, square, location, height, capacity) VALUES ('10', '320', 'Центральная зона', '6.5', '2');
INSERT INTO enclosure (id, square, location, height, capacity) VALUES ('11', '380', 'Центральная зона', '7.5', '3');
INSERT INTO enclosure (id, square, location, height, capacity) VALUES ('12', '420', 'Северо-Восточная зона', '8', '4');
INSERT INTO enclosure (id, square, location, height, capacity) VALUES ('13', '470', 'Юго-Западная зона', '9', '5');
INSERT INTO enclosure (id, square, location, height, capacity) VALUES ('16', '620', 'Юго-Восточная зона', '12.5', '8');
INSERT INTO enclosure (id, square, location, height, capacity) VALUES ('17', '670', 'Северо-Западная зона', '13.5', '9');
INSERT INTO enclosure (id, square, location, height, capacity) VALUES ('18', '720', 'Центральная зона', '14', '10');
INSERT INTO enclosure (id, square, location, height, capacity) VALUES ('19', '770', 'Центральная зона', '15', '11');
INSERT INTO enclosure (id, square, location, height, capacity) VALUES ('20', '800', 'Северная зона', '16', '12');
INSERT INTO enclosure (id, square, location, height, capacity) VALUES ('22', '900', 'Восточная зона', '18', '14');
INSERT INTO enclosure (id, square, location, height, capacity) VALUES ('23', '950', 'Западная зона', '19', '15');
INSERT INTO enclosure (id, square, location, height, capacity) VALUES ('24', '1000', 'Центральная зона', '20', '16');
INSERT INTO enclosure (id, square, location, height, capacity) VALUES ('25', '1050', 'Юго-Западная зона', '21', '17');
INSERT INTO enclosure (id, square, location, height, capacity) VALUES ('26', '1100', 'Северо-Западная зона', '22', '18');
INSERT INTO enclosure (id, square, location, height, capacity) VALUES ('27', '1150', 'Юго-Восточная зона', '23', '19');
INSERT INTO enclosure (id, square, location, height, capacity) VALUES ('14', '530', 'Северная зона', '10.5', '6');
INSERT INTO enclosure (id, square, location, height, capacity) VALUES ('15', '580', 'Южная зона', '11.5', '7');
INSERT INTO enclosure (id, square, location, height, capacity) VALUES ('1', '600', 'Северная зона', '10', '5');
CREATE UNIQUE INDEX enclosure_pkey ON public.enclosure USING btree (id);


--
-- Table: enclosure_employee
--

CREATE TABLE enclosure_employee (enclosure_id integer NOT NULL, employee_id integer NOT NULL);

INSERT INTO enclosure_employee (enclosure_id, employee_id) VALUES ('1', '19');
INSERT INTO enclosure_employee (enclosure_id, employee_id) VALUES ('2', '5');
INSERT INTO enclosure_employee (enclosure_id, employee_id) VALUES ('3', '22');
INSERT INTO enclosure_employee (enclosure_id, employee_id) VALUES ('4', '14');
INSERT INTO enclosure_employee (enclosure_id, employee_id) VALUES ('5', '27');
INSERT INTO enclosure_employee (enclosure_id, employee_id) VALUES ('6', '10');
INSERT INTO enclosure_employee (enclosure_id, employee_id) VALUES ('7', '18');
INSERT INTO enclosure_employee (enclosure_id, employee_id) VALUES ('8', '3');
INSERT INTO enclosure_employee (enclosure_id, employee_id) VALUES ('9', '11');
INSERT INTO enclosure_employee (enclosure_id, employee_id) VALUES ('10', '24');
INSERT INTO enclosure_employee (enclosure_id, employee_id) VALUES ('11', '1');
INSERT INTO enclosure_employee (enclosure_id, employee_id) VALUES ('12', '16');
INSERT INTO enclosure_employee (enclosure_id, employee_id) VALUES ('13', '25');
INSERT INTO enclosure_employee (enclosure_id, employee_id) VALUES ('14', '6');
INSERT INTO enclosure_employee (enclosure_id, employee_id) VALUES ('15', '12');
INSERT INTO enclosure_employee (enclosure_id, employee_id) VALUES ('16', '23');
INSERT INTO enclosure_employee (enclosure_id, employee_id) VALUES ('17', '9');
INSERT INTO enclosure_employee (enclosure_id, employee_id) VALUES ('18', '15');
INSERT INTO enclosure_employee (enclosure_id, employee_id) VALUES ('19', '7');
INSERT INTO enclosure_employee (enclosure_id, employee_id) VALUES ('20', '20');
INSERT INTO enclosure_employee (enclosure_id, employee_id) VALUES ('21', '4');
INSERT INTO enclosure_employee (enclosure_id, employee_id) VALUES ('22', '2');
INSERT INTO enclosure_employee (enclosure_id, employee_id) VALUES ('23', '17');
INSERT INTO enclosure_employee (enclosure_id, employee_id) VALUES ('24', '8');
INSERT INTO enclosure_employee (enclosure_id, employee_id) VALUES ('25', '13');
INSERT INTO enclosure_employee (enclosure_id, employee_id) VALUES ('26', '21');
INSERT INTO enclosure_employee (enclosure_id, employee_id) VALUES ('27', '26');


--
-- Table: exhibit
--

CREATE TABLE exhibit (id integer NOT NULL, name character varying NOT NULL, description text, type character varying, date date);

INSERT INTO exhibit (id, name, description, type, date) VALUES ('1', 'Окаменелость трилобита', 'Окаменелость древнего морского членистоногого, возрастом около 500 миллионов лет.', 'Палеонтология', '2023-01-01');
INSERT INTO exhibit (id, name, description, type, date) VALUES ('5', 'Скелет тираннозавра', 'Полный скелет хищного динозавра, одного из самых известных представителей мелового периода.', 'Палеонтология', '2023-05-20');
INSERT INTO exhibit (id, name, description, type, date) VALUES ('17', 'Окаменелость ихтиозавра', 'Окаменелость древнего морского рептилии, похожего на дельфина.', 'Палеонтология', '2024-05-30');
INSERT INTO exhibit (id, name, description, type, date) VALUES ('20', 'Окаменелость птеродактиля', 'Окаменелость крылатого динозавра, жившего в мезозойскую эру.', 'Палеонтология', '2024-08-15');
INSERT INTO exhibit (id, name, description, type, date) VALUES ('24', 'Окаменелость динозавра-велоцираптора', 'Окаменелость небольшого хищного динозавра.', 'Палеонтология', '2024-12-05');
INSERT INTO exhibit (id, name, description, type, date) VALUES ('26', 'Окаменелость плезиозавра', 'Окаменелость древнего морского рептилии с длинной шеей.', 'Палеонтология', '2025-02-15');
INSERT INTO exhibit (id, name, description, type, date) VALUES ('6', 'Окаменелость анкилозавра', 'Окаменелость бронированного динозавра с мощным хвостом, обитавшего в меловом периоде.', 'Палеонтология', '2023-06-12');
INSERT INTO exhibit (id, name, description, type, date) VALUES ('3', 'Окаменелость стегозавра', 'Окаменелость динозавра с костяными пластинами на спине, жившего в юрском периоде.', 'Палеонтология', '2023-03-10');
INSERT INTO exhibit (id, name, description, type, date) VALUES ('4', 'Окаменелость птеродактиля', 'Окаменелость летающего ящера, обитавшего в мезозойскую эру.', 'Палеонтология', '2023-04-05');
INSERT INTO exhibit (id, name, description, type, date) VALUES ('7', 'Окаменелость ихтиозавра', 'Окаменелость морского рептилии, похожего на современных дельфинов.', 'Палеонтология', '2023-07-18');
INSERT INTO exhibit (id, name, description, type, date) VALUES ('8', 'Окаменелость плезиозавра', 'Окаменелость древнего морского хищника с длинной шеей.', 'Палеонтология', '2023-08-22');
INSERT INTO exhibit (id, name, description, type, date) VALUES ('10', 'Окаменелость диплодока', 'Окаменелость одного из самых длинных динозавров, обитавшего в юрском периоде.', 'Палеонтология', '2023-10-14');
INSERT INTO exhibit (id, name, description, type, date) VALUES ('11', 'Окаменелость спинозавра', 'Окаменелость крупного хищника с парусом на спине, обитавшего в меловом периоде.', 'Палеонтология', '2023-11-05');
INSERT INTO exhibit (id, name, description, type, date) VALUES ('13', 'Окаменелость цератозавра', 'Окаменелость хищного динозавра с рогом на носу.', 'Палеонтология', '2024-01-10');
INSERT INTO exhibit (id, name, description, type, date) VALUES ('15', 'Окаменелость гадрозавра', 'Окаменелость утконосого динозавра, обитавшего в меловом периоде.', 'Палеонтология', '2024-03-20');
INSERT INTO exhibit (id, name, description, type, date) VALUES ('14', 'Окаменелость древнего коралла', 'Окаменелость коралла, жившего в мезозойскую эру, когда формировались современные океаны.', 'Геология', '2024-02-15');
INSERT INTO exhibit (id, name, description, type, date) VALUES ('12', 'Реконструкция экосистемы мезозоя', 'Научная реконструкция экосистемы мезозойской эры, включая флору и фауну.', 'История Земли', '2023-12-01');
INSERT INTO exhibit (id, name, description, type, date) VALUES ('22', 'Реконструкция климата мезозоя', 'Научная реконструкция климатических условий мезозойской эры, когда доминировали динозавры.', 'История Земли', '2024-10-25');
INSERT INTO exhibit (id, name, description, type, date) VALUES ('2', 'Окаменелость вулканической породы', 'Образец вулканической породы, сформировавшейся в мезозойскую эру.', 'Геология и физика', '2023-02-15');
INSERT INTO exhibit (id, name, description, type, date) VALUES ('9', 'Окаменелость древнего леса', 'Окаменелость фрагмента древнего леса, существовавшего в мезозойскую эру.', 'Геология', '2023-09-30');
INSERT INTO exhibit (id, name, description, type, date) VALUES ('16', 'Окаменелость дейнониха', 'Окаменелость хищного динозавра, известного своими серповидными когтями.', 'Палеонтология', '2024-04-25');
INSERT INTO exhibit (id, name, description, type, date) VALUES ('19', 'Окаменелость овираптора', 'Окаменелость динозавра, известного своими гнездами и заботой о потомстве.', 'Палеонтология', '2024-07-10');
INSERT INTO exhibit (id, name, description, type, date) VALUES ('21', 'Окаменелость карнотавра', 'Окаменелость хищного динозавра с рогами на голове.', 'Палеонтология', '2024-09-20');
INSERT INTO exhibit (id, name, description, type, date) VALUES ('23', 'Окаменелость майазавра', 'Окаменелость динозавра, известного своими гнездами и заботой о детенышах.', 'Палеонтология', '2024-11-30');
INSERT INTO exhibit (id, name, description, type, date) VALUES ('25', 'Окаменелость стиракозавра', 'Окаменелость рогатого динозавра с костяным воротником.', 'Палеонтология', '2025-01-10');
INSERT INTO exhibit (id, name, description, type, date) VALUES ('27', 'Окаменелость теропода', 'Окаменелость хищного динозавра, предка современных птиц.', 'Палеонтология', '2025-03-20');
INSERT INTO exhibit (id, name, description, type, date) VALUES ('18', 'Карта мезозойской эры', 'Реконструкция карты Земли в период мезозойской эры, когда доминировали динозавры.', 'История Земли', '2024-06-05');
CREATE UNIQUE INDEX exhibit_pkey ON public.exhibit USING btree (id);


--
-- Table: food
--

CREATE TABLE food (id integer NOT NULL, price money NOT NULL, species character varying, composition text NOT NULL, amount real);

INSERT INTO food (id, price, species, composition, amount) VALUES ('3', '7,15 Br', 'Говядина', 'Мясной', '55.3');
INSERT INTO food (id, price, species, composition, amount) VALUES ('4', '9,99 Br', 'Рыба', 'Белковый', '60');
INSERT INTO food (id, price, species, composition, amount) VALUES ('5', '5,50 Br', 'Овощи, фрукты', 'Растительный', '75.8');
INSERT INTO food (id, price, species, composition, amount) VALUES ('6', '10,25 Br', 'Кости, хрящи', 'Минеральный', '45.6');
INSERT INTO food (id, price, species, composition, amount) VALUES ('7', '8,75 Br', 'Зерновые смеси', 'Злаковый', '85.2');
INSERT INTO food (id, price, species, composition, amount) VALUES ('8', '6,90 Br', 'Курица, субпродукты', 'Комбинированный', '50.4');
INSERT INTO food (id, price, species, composition, amount) VALUES ('9', '12,30 Br', 'Морские водоросли', 'Растительный', '30.7');
INSERT INTO food (id, price, species, composition, amount) VALUES ('10', '7,80 Br', 'Ягоды, орехи', 'Фруктовый', '65.9');
INSERT INTO food (id, price, species, composition, amount) VALUES ('11', '9,45 Br', 'Кролик', 'Мясной', '40.2');
INSERT INTO food (id, price, species, composition, amount) VALUES ('12', '11,10 Br', 'Моллюски, ракообразные', 'Белковый', '55.8');
INSERT INTO food (id, price, species, composition, amount) VALUES ('13', '6,60 Br', 'Травы, листья', 'Растительный', '90.1');
INSERT INTO food (id, price, species, composition, amount) VALUES ('14', '8,20 Br', 'Яйца, насекомые', 'Белковый', '48.3');
INSERT INTO food (id, price, species, composition, amount) VALUES ('15', '10,75 Br', 'Семена, злаки', 'Злаковый', '70.5');
INSERT INTO food (id, price, species, composition, amount) VALUES ('16', '7,30 Br', 'Оленина', 'Мясной', '58.7');
INSERT INTO food (id, price, species, composition, amount) VALUES ('17', '9,85 Br', 'Морковь, тыква', 'Овощной', '62.4');
INSERT INTO food (id, price, species, composition, amount) VALUES ('18', '5,95 Br', 'Костная мука', 'Минеральный', '35.9');
INSERT INTO food (id, price, species, composition, amount) VALUES ('19', '8,50 Br', 'Фруктовые смеси', 'Фруктовый', '68.2');
INSERT INTO food (id, price, species, composition, amount) VALUES ('20', '11,40 Br', 'Рыбные отходы', 'Белковый', '42.6');
INSERT INTO food (id, price, species, composition, amount) VALUES ('21', '6,80 Br', 'Травяные гранулы', 'Растительный', '77.3');
INSERT INTO food (id, price, species, composition, amount) VALUES ('22', '9,10 Br', 'Субпродукты птицы', 'Мясной', '53.1');
INSERT INTO food (id, price, species, composition, amount) VALUES ('23', '12,50 Br', 'Морские ежи, креветки', 'Белковый', '37.8');
INSERT INTO food (id, price, species, composition, amount) VALUES ('24', '7,45 Br', 'Зерно, отруби', 'Злаковый', '80.9');
INSERT INTO food (id, price, species, composition, amount) VALUES ('25', '10,90 Br', 'Фруктовые пюре', 'Фруктовый', '59.4');
INSERT INTO food (id, price, species, composition, amount) VALUES ('26', '8,65 Br', 'Кости, мясо', 'Комбинированный', '46.7');
INSERT INTO food (id, price, species, composition, amount) VALUES ('27', '6,25 Br', 'Листья, побеги', 'Растительный', '88');
INSERT INTO food (id, price, species, composition, amount) VALUES ('1', '6,23 Br', 'Злаки, желуди', 'Растительный', '40.7');
INSERT INTO food (id, price, species, composition, amount) VALUES ('2', '8,50 Br', 'Свинина', 'Растительный', '81.4');
CREATE UNIQUE INDEX food_pkey ON public.food USING btree (id);


--
-- Table: food_dinosaur
--

CREATE TABLE food_dinosaur (food_id integer NOT NULL, dinosaur_id integer NOT NULL);

INSERT INTO food_dinosaur (food_id, dinosaur_id) VALUES ('1', '5');
INSERT INTO food_dinosaur (food_id, dinosaur_id) VALUES ('2', '12');
INSERT INTO food_dinosaur (food_id, dinosaur_id) VALUES ('3', '7');
INSERT INTO food_dinosaur (food_id, dinosaur_id) VALUES ('4', '20');
INSERT INTO food_dinosaur (food_id, dinosaur_id) VALUES ('5', '11');
INSERT INTO food_dinosaur (food_id, dinosaur_id) VALUES ('6', '18');
INSERT INTO food_dinosaur (food_id, dinosaur_id) VALUES ('7', '3');
INSERT INTO food_dinosaur (food_id, dinosaur_id) VALUES ('8', '14');
INSERT INTO food_dinosaur (food_id, dinosaur_id) VALUES ('9', '9');
INSERT INTO food_dinosaur (food_id, dinosaur_id) VALUES ('10', '15');
INSERT INTO food_dinosaur (food_id, dinosaur_id) VALUES ('11', '22');
INSERT INTO food_dinosaur (food_id, dinosaur_id) VALUES ('12', '1');
INSERT INTO food_dinosaur (food_id, dinosaur_id) VALUES ('13', '27');
INSERT INTO food_dinosaur (food_id, dinosaur_id) VALUES ('14', '4');
INSERT INTO food_dinosaur (food_id, dinosaur_id) VALUES ('15', '6');
INSERT INTO food_dinosaur (food_id, dinosaur_id) VALUES ('16', '17');
INSERT INTO food_dinosaur (food_id, dinosaur_id) VALUES ('17', '10');
INSERT INTO food_dinosaur (food_id, dinosaur_id) VALUES ('18', '19');
INSERT INTO food_dinosaur (food_id, dinosaur_id) VALUES ('19', '2');
INSERT INTO food_dinosaur (food_id, dinosaur_id) VALUES ('20', '25');
INSERT INTO food_dinosaur (food_id, dinosaur_id) VALUES ('21', '16');
INSERT INTO food_dinosaur (food_id, dinosaur_id) VALUES ('22', '8');
INSERT INTO food_dinosaur (food_id, dinosaur_id) VALUES ('23', '13');
INSERT INTO food_dinosaur (food_id, dinosaur_id) VALUES ('24', '21');
INSERT INTO food_dinosaur (food_id, dinosaur_id) VALUES ('25', '23');
INSERT INTO food_dinosaur (food_id, dinosaur_id) VALUES ('26', '26');
INSERT INTO food_dinosaur (food_id, dinosaur_id) VALUES ('27', '24');


--
-- Table: tour
--

CREATE TABLE tour (id integer NOT NULL, date date NOT NULL, time time without time zone NOT NULL, duration time without time zone, price money NOT NULL, visitor_id integer);

INSERT INTO tour (id, date, time, duration, price, visitor_id) VALUES ('1', '2025-03-06', '12:00:00', '0000-01-01T01:00:00Z', '44,61 Br', '11');
INSERT INTO tour (id, date, time, duration, price, visitor_id) VALUES ('2', '2025-03-13', '14:00:00', '0000-01-01T02:00:00Z', '41,87 Br', '15');
INSERT INTO tour (id, date, time, duration, price, visitor_id) VALUES ('3', '2025-03-13', '08:00:00', '0000-01-01T02:00:00Z', '48,61 Br', '18');
INSERT INTO tour (id, date, time, duration, price, visitor_id) VALUES ('4', '2025-03-11', '12:00:00', '0000-01-01T01:00:00Z', '39,58 Br', '20');
INSERT INTO tour (id, date, time, duration, price, visitor_id) VALUES ('5', '2025-03-16', '14:00:00', '0000-01-01T01:00:00Z', '24,04 Br', '3');
INSERT INTO tour (id, date, time, duration, price, visitor_id) VALUES ('6', '2025-03-19', '11:00:00', '0000-01-01T02:00:00Z', '23,77 Br', '24');
INSERT INTO tour (id, date, time, duration, price, visitor_id) VALUES ('7', '2025-03-20', '10:00:00', '0000-01-01T02:00:00Z', '26,14 Br', '18');
INSERT INTO tour (id, date, time, duration, price, visitor_id) VALUES ('8', '2025-03-06', '11:00:00', '0000-01-01T01:00:00Z', '44,08 Br', '13');
INSERT INTO tour (id, date, time, duration, price, visitor_id) VALUES ('9', '2025-02-26', '08:00:00', '0000-01-01T02:00:00Z', '35,14 Br', '1');
INSERT INTO tour (id, date, time, duration, price, visitor_id) VALUES ('10', '2025-03-24', '13:00:00', '0000-01-01T02:00:00Z', '33,35 Br', '18');
INSERT INTO tour (id, date, time, duration, price, visitor_id) VALUES ('11', '2025-03-02', '13:00:00', '0000-01-01T02:00:00Z', '29,09 Br', '13');
INSERT INTO tour (id, date, time, duration, price, visitor_id) VALUES ('12', '2025-02-26', '09:00:00', '0000-01-01T02:00:00Z', '27,75 Br', '11');
INSERT INTO tour (id, date, time, duration, price, visitor_id) VALUES ('13', '2025-03-15', '15:00:00', '0000-01-01T01:00:00Z', '39,90 Br', '17');
INSERT INTO tour (id, date, time, duration, price, visitor_id) VALUES ('14', '2025-03-11', '09:00:00', '0000-01-01T01:00:00Z', '23,22 Br', '1');
INSERT INTO tour (id, date, time, duration, price, visitor_id) VALUES ('15', '2025-02-27', '14:00:00', '0000-01-01T02:00:00Z', '20,59 Br', '12');
INSERT INTO tour (id, date, time, duration, price, visitor_id) VALUES ('16', '2025-03-08', '10:00:00', '0000-01-01T01:00:00Z', '26,08 Br', '19');
INSERT INTO tour (id, date, time, duration, price, visitor_id) VALUES ('17', '2025-02-25', '10:00:00', '0000-01-01T01:00:00Z', '28,18 Br', '17');
INSERT INTO tour (id, date, time, duration, price, visitor_id) VALUES ('18', '2025-03-08', '15:00:00', '0000-01-01T02:00:00Z', '24,71 Br', '6');
INSERT INTO tour (id, date, time, duration, price, visitor_id) VALUES ('19', '2025-03-09', '08:00:00', '0000-01-01T01:00:00Z', '25,07 Br', '21');
INSERT INTO tour (id, date, time, duration, price, visitor_id) VALUES ('20', '2025-03-08', '08:00:00', '0000-01-01T02:00:00Z', '25,21 Br', '5');
INSERT INTO tour (id, date, time, duration, price, visitor_id) VALUES ('21', '2025-03-12', '11:00:00', '0000-01-01T02:00:00Z', '47,84 Br', '8');
INSERT INTO tour (id, date, time, duration, price, visitor_id) VALUES ('22', '2025-03-22', '13:00:00', '0000-01-01T02:00:00Z', '34,90 Br', '22');
INSERT INTO tour (id, date, time, duration, price, visitor_id) VALUES ('23', '2025-02-25', '14:00:00', '0000-01-01T02:00:00Z', '44,94 Br', '24');
INSERT INTO tour (id, date, time, duration, price, visitor_id) VALUES ('24', '2025-03-20', '09:00:00', '0000-01-01T02:00:00Z', '45,90 Br', '27');
INSERT INTO tour (id, date, time, duration, price, visitor_id) VALUES ('25', '2025-02-26', '11:00:00', '0000-01-01T01:00:00Z', '31,34 Br', '18');
INSERT INTO tour (id, date, time, duration, price, visitor_id) VALUES ('26', '2025-03-03', '12:00:00', '0000-01-01T02:00:00Z', '47,94 Br', '17');
INSERT INTO tour (id, date, time, duration, price, visitor_id) VALUES ('27', '2025-03-18', '08:00:00', '0000-01-01T01:00:00Z', '44,92 Br', '17');
CREATE UNIQUE INDEX tour_pkey ON public.tour USING btree (id);


--
-- Table: tour_enclosure
--

CREATE TABLE tour_enclosure (tour_id integer NOT NULL, enclosure_id integer NOT NULL);

INSERT INTO tour_enclosure (tour_id, enclosure_id) VALUES ('1', '15');
INSERT INTO tour_enclosure (tour_id, enclosure_id) VALUES ('2', '7');
INSERT INTO tour_enclosure (tour_id, enclosure_id) VALUES ('3', '23');
INSERT INTO tour_enclosure (tour_id, enclosure_id) VALUES ('4', '11');
INSERT INTO tour_enclosure (tour_id, enclosure_id) VALUES ('5', '18');
INSERT INTO tour_enclosure (tour_id, enclosure_id) VALUES ('6', '4');
INSERT INTO tour_enclosure (tour_id, enclosure_id) VALUES ('7', '20');
INSERT INTO tour_enclosure (tour_id, enclosure_id) VALUES ('8', '2');
INSERT INTO tour_enclosure (tour_id, enclosure_id) VALUES ('9', '25');
INSERT INTO tour_enclosure (tour_id, enclosure_id) VALUES ('10', '13');
INSERT INTO tour_enclosure (tour_id, enclosure_id) VALUES ('11', '1');
INSERT INTO tour_enclosure (tour_id, enclosure_id) VALUES ('12', '21');
INSERT INTO tour_enclosure (tour_id, enclosure_id) VALUES ('13', '10');
INSERT INTO tour_enclosure (tour_id, enclosure_id) VALUES ('14', '16');
INSERT INTO tour_enclosure (tour_id, enclosure_id) VALUES ('15', '19');
INSERT INTO tour_enclosure (tour_id, enclosure_id) VALUES ('16', '5');
INSERT INTO tour_enclosure (tour_id, enclosure_id) VALUES ('17', '26');
INSERT INTO tour_enclosure (tour_id, enclosure_id) VALUES ('18', '3');
INSERT INTO tour_enclosure (tour_id, enclosure_id) VALUES ('19', '12');
INSERT INTO tour_enclosure (tour_id, enclosure_id) VALUES ('20', '24');
INSERT INTO tour_enclosure (tour_id, enclosure_id) VALUES ('21', '9');
INSERT INTO tour_enclosure (tour_id, enclosure_id) VALUES ('22', '14');
INSERT INTO tour_enclosure (tour_id, enclosure_id) VALUES ('23', '8');
INSERT INTO tour_enclosure (tour_id, enclosure_id) VALUES ('24', '27');
INSERT INTO tour_enclosure (tour_id, enclosure_id) VALUES ('25', '6');
INSERT INTO tour_enclosure (tour_id, enclosure_id) VALUES ('26', '22');
INSERT INTO tour_enclosure (tour_id, enclosure_id) VALUES ('27', '17');


--
-- Table: tour_exhibit
--

CREATE TABLE tour_exhibit (tour_id integer NOT NULL, exhibit_id integer NOT NULL);

INSERT INTO tour_exhibit (tour_id, exhibit_id) VALUES ('1', '9');
INSERT INTO tour_exhibit (tour_id, exhibit_id) VALUES ('2', '3');
INSERT INTO tour_exhibit (tour_id, exhibit_id) VALUES ('3', '16');
INSERT INTO tour_exhibit (tour_id, exhibit_id) VALUES ('4', '22');
INSERT INTO tour_exhibit (tour_id, exhibit_id) VALUES ('5', '11');
INSERT INTO tour_exhibit (tour_id, exhibit_id) VALUES ('6', '25');
INSERT INTO tour_exhibit (tour_id, exhibit_id) VALUES ('7', '4');
INSERT INTO tour_exhibit (tour_id, exhibit_id) VALUES ('8', '19');
INSERT INTO tour_exhibit (tour_id, exhibit_id) VALUES ('9', '12');
INSERT INTO tour_exhibit (tour_id, exhibit_id) VALUES ('10', '7');
INSERT INTO tour_exhibit (tour_id, exhibit_id) VALUES ('11', '2');
INSERT INTO tour_exhibit (tour_id, exhibit_id) VALUES ('12', '20');
INSERT INTO tour_exhibit (tour_id, exhibit_id) VALUES ('13', '14');
INSERT INTO tour_exhibit (tour_id, exhibit_id) VALUES ('14', '18');
INSERT INTO tour_exhibit (tour_id, exhibit_id) VALUES ('15', '17');
INSERT INTO tour_exhibit (tour_id, exhibit_id) VALUES ('16', '6');
INSERT INTO tour_exhibit (tour_id, exhibit_id) VALUES ('17', '24');
INSERT INTO tour_exhibit (tour_id, exhibit_id) VALUES ('18', '1');
INSERT INTO tour_exhibit (tour_id, exhibit_id) VALUES ('19', '8');
INSERT INTO tour_exhibit (tour_id, exhibit_id) VALUES ('20', '15');
INSERT INTO tour_exhibit (tour_id, exhibit_id) VALUES ('21', '10');
INSERT INTO tour_exhibit (tour_id, exhibit_id) VALUES ('22', '26');
INSERT INTO tour_exhibit (tour_id, exhibit_id) VALUES ('23', '21');
INSERT INTO tour_exhibit (tour_id, exhibit_id) VALUES ('24', '13');
INSERT INTO tour_exhibit (tour_id, exhibit_id) VALUES ('25', '5');
INSERT INTO tour_exhibit (tour_id, exhibit_id) VALUES ('26', '27');
INSERT INTO tour_exhibit (tour_id, exhibit_id) VALUES ('27', '23');


--
-- Table: visitor
--

CREATE TABLE visitor (id integer NOT NULL, name character varying NOT NULL, gender character varying NOT NULL, age integer, ticket integer NOT NULL);

INSERT INTO visitor (id, name, gender, age, ticket) VALUES ('2', 'Мария Петрова', 'Женский', '30', '12346');
INSERT INTO visitor (id, name, gender, age, ticket) VALUES ('4', 'Анна Кузнецова', 'Женский', '28', '12348');
INSERT INTO visitor (id, name, gender, age, ticket) VALUES ('5', 'Сергей Попов', 'Мужской', '35', '12349');
INSERT INTO visitor (id, name, gender, age, ticket) VALUES ('6', 'Елена Соловьёва', 'Женский', '27', '12350');
INSERT INTO visitor (id, name, gender, age, ticket) VALUES ('7', 'Дмитрий Васильев', 'Мужской', '33', '12351');
INSERT INTO visitor (id, name, gender, age, ticket) VALUES ('8', 'Ольга Павлова', 'Женский', '24', '12352');
INSERT INTO visitor (id, name, gender, age, ticket) VALUES ('9', 'Николай Фёдоров', 'Мужской', '29', '12353');
INSERT INTO visitor (id, name, gender, age, ticket) VALUES ('10', 'Татьяна Романова', 'Женский', '31', '12354');
INSERT INTO visitor (id, name, gender, age, ticket) VALUES ('11', 'Владимир Николаев', 'Мужской', '40', '12355');
INSERT INTO visitor (id, name, gender, age, ticket) VALUES ('12', 'Юлия Ковалёва', 'Женский', '26', '12356');
INSERT INTO visitor (id, name, gender, age, ticket) VALUES ('13', 'Максим Егоров', 'Мужской', '21', '12357');
INSERT INTO visitor (id, name, gender, age, ticket) VALUES ('15', 'Андрей Сидоров', 'Мужской', '38', '12359');
INSERT INTO visitor (id, name, gender, age, ticket) VALUES ('16', 'Ксения Лебедева', 'Женский', '23', '12360');
INSERT INTO visitor (id, name, gender, age, ticket) VALUES ('17', 'Роман Григорьев', 'Мужской', '34', '12361');
INSERT INTO visitor (id, name, gender, age, ticket) VALUES ('18', 'Виктория Тихонова', 'Женский', '29', '12362');
INSERT INTO visitor (id, name, gender, age, ticket) VALUES ('19', 'Артур Захаров', 'Мужской', '36', '12363');
INSERT INTO visitor (id, name, gender, age, ticket) VALUES ('20', 'Светлана Михайлова', 'Женский', '30', '12364');
INSERT INTO visitor (id, name, gender, age, ticket) VALUES ('21', 'Егор Сафонов', 'Мужской', '27', '12365');
INSERT INTO visitor (id, name, gender, age, ticket) VALUES ('22', 'Лариса Смирнова', 'Женский', '25', '12366');
INSERT INTO visitor (id, name, gender, age, ticket) VALUES ('23', 'Денис Фролов', 'Мужской', '39', '12367');
INSERT INTO visitor (id, name, gender, age, ticket) VALUES ('24', 'Наталья Морозова', 'Женский', '22', '12368');
INSERT INTO visitor (id, name, gender, age, ticket) VALUES ('25', 'Станислав Кузьмин', 'Мужской', '28', '12369');
INSERT INTO visitor (id, name, gender, age, ticket) VALUES ('26', 'Алёна Костина', 'Женский', '31', '12370');
INSERT INTO visitor (id, name, gender, age, ticket) VALUES ('27', 'Павел Громов', 'Мужской', '24', '12371');
INSERT INTO visitor (id, name, gender, age, ticket) VALUES ('1', 'Иван Иванов', 'Мужской', '16', '12345');
INSERT INTO visitor (id, name, gender, age, ticket) VALUES ('14', 'Ирина Белова', 'Женский', '17', '12358');
INSERT INTO visitor (id, name, gender, age, ticket) VALUES ('3', 'Алина Смирнова', 'Женский', '22', '12347');
CREATE UNIQUE INDEX visitor_pkey ON public.visitor USING btree (id);


--
-- Foreign keys
--

ALTER TABLE food_dinosaur ADD CONSTRAINT food_dinosaur_food_id_fkey FOREIGN KEY (food_id) REFERENCES food(id) NOT VALID;
ALTER TABLE enclosure_employee ADD CONSTRAINT enclosure_employee_enclosure_id_fkey FOREIGN KEY (enclosure_id) REFERENCES enclosure(id) NOT VALID;
ALTER TABLE employee_exhibit ADD CONSTRAINT employee_exhibit_exhibit_id_fkey FOREIGN KEY (exhibit_id) REFERENCES exhibit(id) NOT VALID;
ALTER TABLE food_dinosaur ADD CONSTRAINT food_dinosaur_food_id_fkey1 FOREIGN KEY (food_id) REFERENCES food(id) NOT VALID;
ALTER TABLE enclosure_employee ADD CONSTRAINT enclosure_employee_enclosure_id_fkey1 FOREIGN KEY (enclosure_id) REFERENCES enclosure(id) NOT VALID;
ALTER TABLE enclosure_employee ADD CONSTRAINT enclosure_employee_employee_id_fkey FOREIGN KEY (employee_id) REFERENCES employee(id) NOT VALID;
ALTER TABLE employee_exhibit ADD CONSTRAINT employee_exhibit_employee_id_fkey FOREIGN KEY (employee_id) REFERENCES employee(id) NOT VALID;
ALTER TABLE tour_exhibit ADD CONSTRAINT tour_exhibit_exhibit_id_fkey FOREIGN KEY (exhibit_id) REFERENCES exhibit(id) NOT VALID;
ALTER TABLE tour_enclosure ADD CONSTRAINT tour_enclosure_enclosure_id_fkey FOREIGN KEY (enclosure_id) REFERENCES enclosure(id) NOT VALID;
ALTER TABLE employee_exhibit ADD CONSTRAINT employee_exhibit_exhibit_id_fkey1 FOREIGN KEY (exhibit_id) REFERENCES exhibit(id) NOT VALID;
ALTER TABLE tour_exhibit ADD CONSTRAINT tour_exhibit_exhibit_id_fkey1 FOREIGN KEY (exhibit_id) REFERENCES exhibit(id) NOT VALID;
ALTER TABLE tour_enclosure ADD CONSTRAINT tour_enclosure_enclosure_id_fkey1 FOREIGN KEY (enclosure_id) REFERENCES enclosure(id) NOT VALID;
ALTER TABLE food_dinosaur ADD CONSTRAINT food_dinosaur_food_id_fkey2 FOREIGN KEY (food_id) REFERENCES food(id) NOT VALID;
ALTER TABLE enclosure_employee ADD CONSTRAINT enclosure_employee_enclosure_id_fkey2 FOREIGN KEY (enclosure_id) REFERENCES enclosure(id) NOT VALID;
ALTER TABLE enclosure_employee ADD CONSTRAINT enclosure_employee_employee_id_fkey1 FOREIGN KEY (employee_id) REFERENCES employee(id) NOT VALID;
ALTER TABLE employee_exhibit ADD CONSTRAINT employee_exhibit_employee_id_fkey1 FOREIGN KEY (employee_id) REFERENCES employee(id) NOT VALID;
ALTER TABLE employee_exhibit ADD CONSTRAINT employee_exhibit_exhibit_id_fkey2 FOREIGN KEY (exhibit_id) REFERENCES exhibit(id) NOT VALID;
ALTER TABLE tour_exhibit ADD CONSTRAINT tour_exhibit_exhibit_id_fkey2 FOREIGN KEY (exhibit_id) REFERENCES exhibit(id) NOT VALID;
ALTER TABLE tour_enclosure ADD CONSTRAINT tour_enclosure_enclosure_id_fkey2 FOREIGN KEY (enclosure_id) REFERENCES enclosure(id) NOT VALID;
ALTER TABLE food_dinosaur ADD CONSTRAINT food_dinosaur_food_id_fkey3 FOREIGN KEY (food_id) REFERENCES food(id) NOT VALID;
ALTER TABLE enclosure_employee ADD CONSTRAINT enclosure_employee_enclosure_id_fkey3 FOREIGN KEY (enclosure_id) REFERENCES enclosure(id) NOT VALID;
ALTER TABLE enclosure_employee ADD CONSTRAINT enclosure_employee_employee_id_fkey2 FOREIGN KEY (employee_id) REFERENCES employee(id) NOT VALID;
ALTER TABLE employee_exhibit ADD CONSTRAINT employee_exhibit_employee_id_fkey2 FOREIGN KEY (employee_id) REFERENCES employee(id) NOT VALID;
ALTER TABLE employee_exhibit ADD CONSTRAINT employee_exhibit_exhibit_id_fkey3 FOREIGN KEY (exhibit_id) REFERENCES exhibit(id) NOT VALID;
ALTER TABLE tour_exhibit ADD CONSTRAINT tour_exhibit_exhibit_id_fkey3 FOREIGN KEY (exhibit_id) REFERENCES exhibit(id) NOT VALID;
ALTER TABLE tour_enclosure ADD CONSTRAINT tour_enclosure_enclosure_id_fkey3 FOREIGN KEY (enclosure_id) REFERENCES enclosure(id) NOT VALID;
ALTER TABLE dinosaur ADD CONSTRAINT dinosaur_enclosure_id_fkey FOREIGN KEY (enclosure_id) REFERENCES enclosure(id) NOT VALID;
ALTER TABLE tour ADD CONSTRAINT tour_visitor_id_fkey FOREIGN KEY (visitor_id) REFERENCES visitor(id) NOT VALID;
ALTER TABLE food_dinosaur ADD CONSTRAINT food_dinosaur_food_id_fkey4 FOREIGN KEY (food_id) REFERENCES food(id) NOT VALID;
ALTER TABLE food_dinosaur ADD CONSTRAINT food_dinosaur_dinosaur_id_fkey FOREIGN KEY (dinosaur_id) REFERENCES dinosaur(id) NOT VALID;
ALTER TABLE enclosure_employee ADD CONSTRAINT enclosure_employee_enclosure_id_fkey4 FOREIGN KEY (enclosure_id) REFERENCES enclosure(id) NOT VALID;
ALTER TABLE enclosure_employee ADD CONSTRAINT enclosure_employee_employee_id_fkey3 FOREIGN KEY (employee_id) REFERENCES employee(id) NOT VALID;
ALTER TABLE employee_exhibit ADD CONSTRAINT employee_exhibit_employee_id_fkey3 FOREIGN KEY (employee_id) REFERENCES employee(id) NOT VALID;
ALTER TABLE employee_exhibit ADD CONSTRAINT employee_exhibit_exhibit_id_fkey4 FOREIGN KEY (exhibit_id) REFERENCES exhibit(id) NOT VALID;
ALTER TABLE tour_exhibit ADD CONSTRAINT tour_exhibit_tour_id_fkey FOREIGN KEY (tour_id) REFERENCES tour(id) NOT VALID;
ALTER TABLE tour_exhibit ADD CONSTRAINT tour_exhibit_exhibit_id_fkey4 FOREIGN KEY (exhibit_id) REFERENCES exhibit(id) NOT VALID;
ALTER TABLE tour_enclosure ADD CONSTRAINT tour_enclosure_tour_id_fkey FOREIGN KEY (tour_id) REFERENCES tour(id) NOT VALID;
ALTER TABLE tour_enclosure ADD CONSTRAINT tour_enclosure_enclosure_id_fkey4 FOREIGN KEY (enclosure_id) REFERENCES enclosure(id) NOT VALID;
