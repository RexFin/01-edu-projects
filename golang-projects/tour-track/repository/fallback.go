package repository

import (
	"embed"
	_ "embed"
	"encoding/json"
	"fmt"
	"strings"
)

//go:embed preview_dates.json preview_locations.json
var previewFiles embed.FS

type previewDates struct {
	Index []dateData `json:"index"`
}

type previewLocations struct {
	Index []locationData `json:"index"`
}

// FallbackData returns an offline catalog with local artist photos and the
// last available dates and locations snapshot from the source API.
func FallbackData() ([]Artist, map[int]Relation) {
	artists := []Artist{
		{1, image("Queen"), "Queen", []string{"Freddie Mercury", "Brian May", "John Daecon", "Roger Meddows-Taylor", "Mike Grose", "Barry Mitchell", "Doug Fogie"}, 1970, "14-12-1973"},
		{2, image("SOJA"), "SOJA", []string{"Jacob Hemphill", "Bob Jefferson", "Ryan \"Byrd\" Berty", "Ken Brownell", "Patrick O'Shea", "Hellman Escorcia", "Rafael Rodriguez", "Trevor Young"}, 1997, "05-06-2002"},
		{3, image("Pink Floyd"), "Pink Floyd", []string{"David Gilmour", "Syd Barrett", "Roger Waters", "Richard Wright", "Nick Mason", "Bob Klose"}, 1965, "05-08-1967"},
		{4, image("Scorpions"), "Scorpions", []string{"Klaus Meine", "Rudolf Schenker", "Matthias Jabs", "Pawel Maciwoda", "Mikkey Dee"}, 1965, "21-04-1979"},
		{5, image("XXXTentacion"), "XXXTentacion", []string{"Jahseh Dwayne Ricardo Onfroy"}, 2013, "07-09-2017"},
		{6, image("Mac Miller"), "Mac Miller", []string{"Malcolm James McCormick"}, 2007, "29-06-2013"},
		{7, image("Joyner Lucas"), "Joyner Lucas", []string{"Gary Maurice Lucas Jr."}, 2007, "13-05-2011"},
		{8, image("Kendrick Lamar"), "Kendrick Lamar", []string{"Kendrick Lamar Duckworth"}, 2004, "03-11-2012"},
		{9, image("AC/DC"), "AC/DC", []string{"Angus Young", "Chris Slade", "Stevie Young", "Axl Rose"}, 1973, "11-05-1977"},
		{10, image("Pearl Jam"), "Pearl Jam", []string{"Eddie Vedder", "Mike McCready", "Stone Gossard", "Jeff Ament", "Matt Cameron"}, 1990, "03-07-1992"},
		{11, image("Katy Perry"), "Katy Perry", []string{"Katheryn Elizabeth Hudson"}, 2001, "04-10-2008"},
		{12, image("Rihanna"), "Rihanna", []string{"Robyn Rihanna Fenty"}, 2003, "10-09-2005"},
		{13, image("Genesis"), "Genesis", []string{"Mike Rutherford", "Tony Banks", "Phil Collins", "Chester Thompson", "Daryl Stuermer"}, 1967, "14-10-1972"},
		{14, image("Phil Collins"), "Phil Collins", []string{"Phil Collins"}, 1966, "21-02-1981"},
		{15, image("Led Zeppelin"), "Led Zeppelin", []string{"Jimmy Page", "John Paul Jones", "Robert Plant", "John Bonham"}, 1968, "12-04-1969"},
		{16, image("The Jimi Hendrix Experience"), "The Jimi Hendrix Experience", []string{"Jimi Hendrix", "Mitch Mitchell", "Noel Redding", "Billy Cox"}, 1966, "27-05-1967"},
		{17, image("Bee Gees"), "Bee Gees", []string{"Barry Gibb", "Robin Gibb", "Maurice Gibb"}, 1958, "12-08-1967"},
		{18, image("Deep Purple"), "Deep Purple", []string{"Ian Paice", "Ian Gillan", "Roger Glover", "Steve Morse", "Don Airey"}, 1968, "24-01-1970"},
		{19, image("Aerosmith"), "Aerosmith", []string{"Steven Tyler", "Tom Hamilton", "Joey Kramer", "Joe Perry", "Brad Whitford"}, 1970, "05-09-1987"},
		{20, image("Dire Straits"), "Dire Straits", []string{"Mark Knopfler", "David Knopfler", "John Illsley", "Pick Withers"}, 1977, "22-07-1978"},
		{21, image("Mamonas Assassinas"), "Mamonas Assassinas", []string{"Dinho", "Bento Hinoto", "Júlio Rasec", "Sérgio Reoli", "Samuel Reis de Oliveira"}, 1995, "23-06-1995"},
		{22, image("Thirty Seconds to Mars"), "Thirty Seconds to Mars", []string{"Jared Leto", "Shannon Leto", "Tomo Miličević", "Matt Wachter", "Solon Bixler"}, 1998, "27-06-2002"},
		{23, image("Imagine Dragons"), "Imagine Dragons", []string{"Dan Reynolds", "Daniel Wayne Sermon", "Ben McKee"}, 2008, "04-09-2012"},
		{24, image("Juice WRLD"), "Juice WRLD", []string{"Jarad Higgins"}, 2015, "23-05-2018"},
		{25, image("Logic"), "Logic", []string{"Sir Robert Bryson Hall II"}, 2009, "22-09-2015"},
		{26, image("Alec Benjamin"), "Alec Benjamin", []string{"Alec Benjamin"}, 2013, "15-11-2018"},
		{27, image("Bobby McFerrin"), "Bobby McFerrin", []string{"Bobby McFerrin"}, 1977, "01-09-1982"},
		{28, image("R3HAB"), "R3HAB", []string{"Fadil El Ghoul"}, 2008, "09-12-2014"},
		{29, image("Post Malone"), "Post Malone", []string{"Post Malone"}, 2013, "13-05-2016"},
		{30, image("Travis Scott"), "Travis Scott", []string{"Jacques Bermon Webster II"}, 2008, "18-08-2014"},
		{31, image("J. Cole"), "J. Cole", []string{"Jermaine Lamarr Cole"}, 2007, "04-05-2007"},
		{32, image("Nickelback"), "Nickelback", []string{"Chad Kroeger", "Ryan Peake", "Mike Kroeger", "Daniel Adair", "Brandon Kroeger", "Ryan Vikedal", "Mitch Guindon"}, 1995, "04-09-1996"},
		{33, image("Mobb Deep"), "Mobb Deep", []string{"Prodigy", "Havoc"}, 1991, "23-04-1993"},
		{34, image("Guns N' Roses"), "Guns N' Roses", []string{"Axl Rose", "Duff McKagan", "Slash", "Dizzy Reed", "Richard Fortus", "Frank Ferrer", "Melissa Reese"}, 1985, "21-07-1987"},
		{35, image("N.W.A."), "N.W.A", []string{"Ice Cube", "Eazy-E", "Dr. Dre", "MC Ren", "DJ Yella", "The D.O.C.", "Arabian Prince"}, 1986, "03-11-1988"},
		{36, image("U2"), "U2", []string{"Bono", "The Edge", "Adam Clayton", "Larry Mullen Jr.", "Ivan McCormick", "Peter Martin", "Dick Evans"}, 1976, "06-12-1980"},
		{37, image("Arctic Monkeys"), "Arctic Monkeys", []string{"Alex Turner", "Matt Helders", "Jamie Cook", "Nick O'Malley", "Andy Nicholson", "Glyn Jones"}, 2002, "23-04-2007"},
		{38, image("Fall Out Boy"), "Fall Out Boy", []string{"Pete Wentz", "Patrick Stump", "Andy Hurley", "Joe Trohman"}, 2001, "25-03-2003"},
		{39, image("Gorillaz"), "Gorillaz", []string{"2-D", "Noodle", "Murdoc Niccals", "Russel Hobbs"}, 1998, "26-03-2001"},
		{40, image("Eagles"), "Eagles", []string{"Glenn Frey", "Don Henley", "Randy Meisner", "Bernie Leadon"}, 1971, "01-06-1972"},
		{41, image("Linkin Park"), "Linkin Park", []string{"Chester Bennington", "Mike Shinoda", "Brad Delson", "Dave Farrell", "Joe Hahn", "Rob Bourdon"}, 1996, "27-01-2001"},
		{42, image("Red Hot Chili Peppers"), "Red Hot Chili Peppers", []string{"Anthony Kiedis", "Flea", "Chad Smith", "Josh Klinghoffer"}, 1983, "10-02-1990"},
		{43, image("Eminem"), "Eminem", []string{"Marshall Bruce Mathers"}, 1996, "10-04-1999"},
		{44, image("Green Day"), "Green Day", []string{"Billie Joe Armstrong", "Mike Dirnt", "Tré Cool"}, 1986, "18-06-1994"},
		{45, image("Metallica"), "Metallica", []string{"James Hetfield", "Lars Ulrich", "Kirk Hammett", "Robert Trujillo"}, 1981, "22-08-1987"},
		{46, image("Coldplay"), "Coldplay", []string{"Chris Martin", "Guy Berryman", "Jonny Buckland", "Will Champion"}, 1996, "22-07-2000"},
		{47, image("Maroon 5"), "Maroon 5", []string{"Adam Levine", "Jesse Carmichael", "Mickey Madden", "James Valentine", "Matt Flynn", "PJ Morton", "Sam Farrar"}, 1994, "31-01-2004"},
		{48, image("Twenty One Pilots"), "Twenty One Pilots", []string{"Tyler Joseph", "Josh Dun"}, 2009, "14-09-2013"},
		{49, image("The Rolling Stones"), "The Rolling Stones", []string{"Mick Jagger", "Keith Richards", "Charlie Watts", "Ron Wood"}, 1962, "31-07-1963"},
		{50, image("Muse"), "Muse", []string{"Matthew Bellamy", "Christopher Wolstenholme", "Dominic Howard"}, 1994, "26-06-1999"},
		{51, image("Foo Fighters"), "Foo Fighters", []string{"Dave Grohl", "Nate Mendel", "Taylor Hawkins", "Chris Shiflett", "Pat Smear", "Rami Jaffee"}, 1994, "01-07-1995"},
		{52, image("The Chainsmokers"), "The Chainsmokers", []string{"Alexander Pall", "Andrew Taggart", "Matt McGuire", "Tony Ann"}, 2008, "15-03-2014"},
	}

	var dates previewDates
	var locations previewLocations
	if err := decodePreview("preview_dates.json", &dates); err != nil {
		panic(err)
	}
	if err := decodePreview("preview_locations.json", &locations); err != nil {
		panic(err)
	}
	relations := make(map[int]Relation, len(artists))
	for _, item := range dates.Index {
		relation := relations[item.ID]
		relation.ID, relation.Dates = item.ID, item.Dates
		relations[item.ID] = relation
	}
	for _, item := range locations.Index {
		relation := relations[item.ID]
		relation.ID, relation.Locations = item.ID, item.Locations
		relations[item.ID] = relation
	}
	return artists, relations
}

func decodePreview(name string, dst any) error {
	data, err := previewFiles.ReadFile(name)
	if err != nil {
		return fmt.Errorf("read local %s: %w", name, err)
	}
	if err := json.Unmarshal(data, dst); err != nil {
		return fmt.Errorf("decode local %s: %w", name, err)
	}
	return nil
}

func image(name string) string {
	var slug strings.Builder
	for _, r := range strings.ToLower(name) {
		if r >= 'a' && r <= 'z' || r >= '0' && r <= '9' {
			slug.WriteRune(r)
		} else if slug.Len() > 0 && !strings.HasSuffix(slug.String(), "-") {
			slug.WriteByte('-')
		}
	}
	return "/static/images/artists/" + strings.Trim(slug.String(), "-") + ".jpg"
}
