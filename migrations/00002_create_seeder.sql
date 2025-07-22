-- +goose Up
-- +goose StatementBegin
INSERT INTO categories (name, description)
VALUES
('Beef', 'Diced Beef is always a great meat to have in the freezer to make simple slow cooked recipes like stews, goulash, Beef Skewers etc. The Beef Tri Tip Steak used to make our Diced Beef has no gristle or excess fat '),
('Lamb', 'Expert butchers cut these luxury Lamb Leg Steaks from our special “Tunnel Boned” Legs of Lamb. Your Leg Steaks are approximately 15-20 mm wide through the full leg, each leg of lamb gives 5-7 steaks.'),
('turkey', 'The generous quality Turkey Wings are filling & tasty with generous amounts of meat on the bone. The Wings are great when Smoked on the Barbecue & easy to cook in the Oven or Air Fryer'),
('Chicken', 'In a meat shop, chicken can be categorized by several factors including its grade, age/size, cut, and processing method. Common categories include broiler/fryer chickens, roasters, and Cornish game hens based on age and size.');


INSERT INTO products (name, description, price, category_id)
VALUES
(
    'Diced Beef',
	'meat to have in the freezer to make simple slow cooked recipes like stews, goulash, Beef Skewers etc. The Beef Tri Tip Steak used to make our Diced Beef has no gristle or excess fat ',
	12.45,
	1
),
(
    'Beef T Bone Steaks',
	'T Bone Steaks are cut from our Irish Grass Fed select Beef, these hand cut Steaks are big thick & juicy, perfect for sharing on those special Steak Nites! The T Bone Steaks are actually called ',
	22.15,
	1
),
(
    'Beef Hanger Steak',
	'Sherwood Beef Hanger Steak is a Grass Fed rich flavoured piece, great on the Barbecue or grilled indoors.',
	33.22,
	1
),

(
    'Lamb Leg Steaks',
	'Your Leg Steaks are approximately 15-20 mm wide through the full leg, each leg of lamb gives 5-7 steaks. The Lamb Leg Steaks give great results cooked under the grill or in a hot skillet',
	9.21,
	2
),
(
    'Lamb Shanks',
	'The generous Lamb Shanks are cut by expert butchers are meaty, making them excellent value for money when shopping online. Enjoy the flavour of natural tasting Shanks from finest Lambs grazed on some of the worlds best Green pastures.',
	41.48,
	2
),
(
    'Lamb Loin Chops',
	'Cutting the Lamb Loin Chops from the Saddle or Loin of Lamb give the best value for money because they are the same as T-Bone Steaks; with Sirloin & Fillet on the bone. The texture of the Sirloin is medium to fine & the texture of the Lamb Fillet is fine',
	76.32,
	2
),

(
    'Boneless Turkey Breast',
	'We offer a delicious home delivered Boneless Turkey Breast with succulent natural flavoured meat & good texture. The Boneless Breasts are from good quality Turkeys, cut, prepared',
	51.65,
	3
),
(
    'Turkey Wings',
	'The generous quality Turkey Wings are filling & tasty with generous amounts of meat on the bone. The Wings are great when Smoked on the Barbecue & easy to cook in the Oven or Air Fryer.',
	29.31,
	3
),
(
    'Turkey Drumsticks', 
	'The delicious flavour of the darker meat on the Turkey Drumsticks gets best results when cooked at a lower temperature for longer. There is a generous amount of tender succulent meat on the bone,',
	34.87,
	3
),

(
    'Chicken Meat Box', 
	'The delicious flavour of the darker meat on the Turkey Drumsticks gets best results when cooked at a lower temperature for longer. There is a generous amount of tender succulent meat on the bone,',
	71.26,
	4
),
(
    'Chicken Wings', 
	'Chicken Wings are from our Grade A locally sourced Chickens to production facilities. The Meat on our Chicken Wings stays moist when cooked on the Barbecue or in the oven.',
	41.83,
	4
),
(
	'Chicken Mid Wings', 
	'The delicious flavour of the darker meat on the Turkey Drumsticks gets best results when cooked at a lower temperature for longer. There is a generous amount of tender succulent meat on the bone,',
	34.87,
	4
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM products;
DELETE FROM categories;
-- +goose StatementEnd
