INSERT INTO gotemp_db.pages (page_id, parent_id, route_id, label, data)
VALUES  (1, NULL, 1, 'MENU', NULL),
        (2, 1, 1, 'logo', 'https://placeholder.co/100x100'),
        (3, 1, 1, 'logo_href', '/'),
        (4, 1, 1, 'main_menu', '[
  {
    "label": "Home",
    "href": "/home"
  },
  {
    "label": "About Us",
    "href": "/about"
  },
  {
    "label": "Services",
    "href": "/services"
  },
  {
    "label": "Portfolio",
    "href": "/portfolio"
  },
  {
    "label": "Blog",
    "href": "/blog"
  },
  {
    "label": "Contact",
    "href": "/contact"
  }
]
');