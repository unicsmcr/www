create table User(
	FirstName varchar(30) not null,
	LastName varchar(30) not null,
	Email varchar(100) not null,
	SubscribedToArticles boolean not null,
	SubscribedToEvents boolean not null,
	primary key (Email))