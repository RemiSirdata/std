package main

type Hotel struct {
	Name        string
	Address     string
	PhoneNumber string
	Photo       string
	Distance    float64
	MapUrl      string
}

func getHotelList() []Hotel {
	return []Hotel{
		{"Château de Tilques", "Rue du Château, TILQUES", "09 70 38 23 17", "/assets/hotels/chateau.jpg", 0.1, "https://www.google.fr/maps/place/Najeti+Hôtel+Château+Tilques/@50.775018,2.2119213,17z"},
		{"Chambres du Marais", "3 Rue du Moulin TILQUES", "03 21 88 89 20", "/assets/hotels/chambre-marais.jpg", 1.8, "https://www.google.fr/maps/place/Les+Chambres+du+Marais/@50.7737522,2.198176,15z"},
		{"VVF", "Rue du Rivage SALPERWICK", "03 21 93 79 40", "/assets/hotels/vvf.jpg", 2.2, "https://www.google.fr/maps/place/VVF+Villages+Le+Marais+Audomarois/@50.772734,2.232906,15z"},
		{"Aux Frangins", "3 rue Carnot SAINT OMER", "03 21 38 12 47", "/assets/hotels/aux-frangins.jpg", 5.4, "https://www.google.fr/maps/place/Hôtel+-+Restaurant+Logis+Les+Frangins/@50.7486212,2.2513599,17z"},
		{"Ibis", "4 rue Henri Dupuis SAINT OMER", "03 21 93 11 11", "/assets/hotels/ibis.jpg", 5.2, "https://www.google.fr/maps/place/4+Rue+Henri+Dupuis,+62500+Saint-Omer/@50.7480532,2.2494718,17z"},
		{"Ibis Budget", "Bd Vauban SAINT OMER", "03 21 93 11 11", "/assets/hotels/ibisbudget.jpg", 5.7, "https://www.google.fr/maps/place/Boulevard+Vauban,+62500+Saint-Omer/@50.7476689,2.248416,17z"},
		{"LE CHIC Ô RAIL", "Place du 8 mai 1945, 62500 SAINT OMER", "03 21 93 59 98", "/assets/hotels/chic.jpg", 6, "https://www.google.fr/maps/place/HOTEL+RESTAURANT+LE+CHIC+Ô+RAIL/@50.753115,2.2649723,17z"},
		{"Saint Louis", "25 rue d'Arras SAINT OMER", "03 21 38 35 21", "/assets/hotels/saintlouis.jpg", 6.2, "https://www.google.fr/maps/place/25+Rue+d'Arras,+62500+Saint-Omer/@50.7465195,2.2553332,17z"},
		{"Gîtes Les Libellules", "62 Bis RN 43 TILQUES", "03 21 98 49 28", "/assets/hotels/libellules.jpg", 7.3, "https://www.google.fr/maps/place/Route+Nationale+43/@50.7952858,2.1292702,17z"},
		{"La Sapinière", "12 rte de Setques WISQUES", "03 21 38 94 00", "/assets/hotels/sapi.jpg", 9, "https://www.google.fr/maps/place/12+Route+de+Setques,+62219+Wisques/@50.7222926,2.1921674,17z"},
		{"Le Bretagne", "2 pl Vainquai SAINT OMER", "03 21 38 25 78", "/assets/hotels/bretagne.jpg", 5.9, "https://www.google.fr/maps/place/2+Place+du+Vainquai,+62500+Saint-Omer/@50.7516435,2.2601614,17z"},
		{"Lemon Hotel", "ZAC du Lobel- N43 ARQUES", "03 21 93 81 20", "/assets/hotels/lemo.jpg", 14.5, "https://www.google.fr/maps/place/HOTEL+DU+GOLF/@50.7151748,2.1335842,12z"},
	}
}
