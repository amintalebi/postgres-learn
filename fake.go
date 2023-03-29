package main

import (
	"math/rand"
	"strconv"
	"time"
)

func GenerateFakeName() (string, string) {
	rand.Seed(time.Now().UnixNano())

	firstNames := []string{
		"Emily", "Emma", "Madison", "Ava", "Olivia",
		"Sophia", "Abigail", "Isabella", "Mia", "Charlotte",
		"Amelia", "Evelyn", "Elizabeth", "Sofia", "Avery",
		"Ella", "Scarlett", "Grace", "Chloe", "Victoria",
		"Aubrey", " Harper", "Ellie", "Adalynn", "Riley",
		"Aria", "Hailey", "Natalie", "Camila", "Hazel",
	}
	lastNames := []string{
		"Smith", "Johnson", "Williams", "Jones", "Brown",
		"Davis", "Miller", "Wilson", "Moore", "Taylor",
		"Anderson", "Thomas", "Jackson", "White", "Harris",
		"Martin", "Thompson", "Young", "Allen", "King",
		"Wright", "Scott", "Green", "Baker", "Adams",
		"Nelson", "Carter", "Mitchell", "Perez", "Roberts",
		"Turner", "Phillips", "Campbell", "Parker", "Evans",
		"Edwards", "Collins", "Stewart", "Sanchez", "Morris",
	}

	return firstNames[rand.Intn(len(firstNames))], lastNames[rand.Intn(len(lastNames))]
}

var firstWords = []string{"The", "A", "An", "My", "Your", "Our", "Their", "His", "Her", "Its"}
var secondWords = []string{"Journey", "Adventure", "Story", "Tale", "Mystery", "Legend", "Epic", "Quest", "Dream", "Destiny"}
var thirdWords = []string{"of the", "in the", "to the", "from the", "with the", "without the", "among the", "through the", "towards the", "by the"}
var fourthWords = []string{
	"Dark", "Bright", "Fierce", "Gentle", "Wild",
	"Tall", "Short", "Thin", "Fat", "Smooth",
	"Rough", "Tough", "Tender", "Tiny", "Giant",
	"Silly", "Happy", "Sad", "Witty", "Serious",
	"Funny", "Loud", "Quiet", "Lively", "Majestic",
	"Radiant", "Gleaming", "Glowing", "Shining", "Luminous",
	"Glorious", "Splendid", "Dazzling", "Brilliant", "Magnificent",
	"Elegant", "Sophisticated", "Refined", "Stylish", "Fashionable",
	"Charming", "Alluring", "Captivating", "Enchanting", "Bewitching",
	"Mystical", "Magical", "Enigmatic", "Mysterious", "Cryptic",
	"Dreamy", "Ethereal", "Otherworldly", "Fantastical", "Unworldly",
	"Fierce", "Fearless", "Bold", "Brave", "Gallant",
	"Sturdy", "Strong", "Robust", "Durable", "Enduring",
	"Graceful", "Effortless", "Smooth", "Fluid", "Agile",
	"Lively", "Energetic", "Vibrant", "Exuberant", "Enthusiastic",
}

var fifthWords = []string{
	"Moon", "Star", "Storm", "Ocean", "Mountain",
	"River", "Tree", "Flower", "Rainbow", "Thunder",
	"Fire", "Dragon", "Phoenix", "Unicorn", "Angel",
	"Demon", "Elf", "Goblin", "Dwarf", "Ogre",
	"Wizard", "Sorcerer", "Warrior", "Knight", "Mage",
	"Kingdom", "Empire", "Realm", "Nation", "State",
	"City", "Town", "Village", "Hamlet", "Fortress",
	"Castle", "Palace", "Manor", "Estate", "Villa",
	"Temple", "Sanctuary", "Chapel", "Shrine", "Cathedral",
	"Mystery", "Secret", "Enigma", "Riddle", "Puzzle",
	"Legend", "Myth", "Fable", "Tale", "Story",
	"Dream", "Vision", "Fantasy", "Imagination", "Reality",
	"Adventure", "Quest", "Journey", "Pilgrimage", "Expedition",
	"Odyssey", "Saga", "Epic", "Chronicle", "Annals"}

func GenerateFakeBookTitle() string {
	rand.Seed(time.Now().UnixNano())
	first := firstWords[rand.Intn(len(firstWords))]
	second := secondWords[rand.Intn(len(secondWords))]
	third := thirdWords[rand.Intn(len(thirdWords))]
	fourth := fourthWords[rand.Intn(len(fourthWords))]
	fifth := fifthWords[rand.Intn(len(fifthWords))]
	return first + " " + second + " " + third + " " + fourth + " " + fifth
}

func GenerateFakePublisher() string {
	publishers := []string{
		"Penguin Random House",
		"Simon & Schuster",
		"HarperCollins",
		"Hachette Livre",
		"Macmillan Publishers",
		"Scholastic Corporation",
		"Wiley",
		"Springer Nature",
		"Elsevier",
		"Taylor & Francis",
		"Bloomsbury Publishing",
		"Blackwell Publishing",
		"Oxford University Press",
		"Cambridge University Press",
		"Routledge",
		"Thames & Hudson",
		"John Benjamins Publishing Company",
		"Edward Elgar Publishing",
		"Emerald Group Publishing",
		"SAGE Publications",
		"Cengage Learning",
		"Pearson Education",
		"McGraw-Hill Education",
		"Houghton Mifflin Harcourt",
		"Dover Publications",
		"MIT Press",
		"Princeton University Press",
		"Yale University Press",
		"Columbia University Press",
		"University of Chicago Press",
		"Harvard Books",
		"Farrar, Straus and Giroux",
		"Random House",
		"Knopf Doubleday Publishing Group",
		"Alfred A. Knopf",
		"Vintage Books",
		"Anchor Books",
		"Ballantine Books",
		"Bantam Books",
		"Del Rey",
		"Spectra",
		"Berkley Publishing Group",
		"Putnam",
		"Riverhead Books",
		"G.P. Putnam's Sons",
		"Roc",
		"Daw Books",
		"Baen Books",
		"Tor Books",
		"Orb Books",
		"Asimov's Science Fiction",
	}

	rand.Seed(time.Now().UnixNano())
	return publishers[rand.Intn(len(publishers))]
}

func GenerateFakeISBN() string {
	rand.Seed(time.Now().UnixNano())
	var fakeISBN string
	for i := 0; i < 13; i++ {
		fakeISBN += strconv.Itoa(rand.Intn(10))
	}
	return fakeISBN
}

func RandomInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Intn(max-min+1)
}

func GenerateFakePhoneNumber() string {
	rand.Seed(time.Now().UnixNano())
	var fakePhoneNumber string
	for i := 0; i < 10; i++ {
		fakePhoneNumber += strconv.Itoa(rand.Intn(10))
	}
	return fakePhoneNumber
}

var streetNames = []string{"Main", "Oak", "Maple", "Pine", "Elm", "Cedar", "Walnut", "Park", "High", "Washington", "Green", "Red", "Yellow", "Blue", "Purple", "Pink", "Brown", "Black", "White", "Gray", "Orange", "River", "Mountain", "Lake", "Hill", "Valley", "Beach", "Forest", "Desert", "Island", "Field", "Glen", "Lake", "View", "River", "Rock", "Meadow", "Field", "Plains", "Coast", "Bluff", "Creek", "Glade", "Lodge", "Ridge", "Run", "Trail", "Way", "Wood"}
var streetTypes = []string{"St", "Ave", "Rd", "Blvd", "Ln", "Dr", "Way", "Pl", "Cir", "Ct"}
var cities = []string{"New York", "Los Angeles", "Chicago", "Houston", "Phoenix", "Philadelphia", "San Antonio", "San Diego", "Dallas", "San Jose", "Austin", "Jacksonville", "Fort Worth", "Columbus", "San Francisco", "Charlotte", "Indianapolis", "Seattle", "Denver", "Washington", "Boston", "El Paso", "Detroit", "Nashville", "Memphis", "Portland", "Oklahoma City", "Las Vegas", "Louisville", "Baltimore", "Milwaukee", "Albuquerque", "Tucson", "Fresno", "Sacramento", "Mesa", "Kansas City", "Atlanta", "Long Beach", "Colorado Springs", "Miami", "Raleigh", "Omaha", "Minneapolis", "New Orleans", "Wichita", "Arlington"}
var states = []string{"NY", "CA", "IL", "TX", "AZ", "PA", "NJ", "FL", "OH", "GA", "MI", "NC", "VA", "MA", "MN", "CO", "IN", "MD", "AL", "MO", "TN", "OR", "WA", "LA", "WI", "UT", "KY", "NM", "IA", "NE", "SC", "CT", "OK", "ND", "MS", "ID", "AR", "ME", "NV", "NH", "KS", "DE", "RI", "HI", "MT", "WV", "WY", "AK", "SD"}

func GenerateFakeAddress() string {
	rand.Seed(time.Now().UnixNano())
	var fakeAddress string
	fakeAddress += strconv.Itoa(rand.Intn(10000)) + " "
	fakeAddress += streetNames[rand.Intn(len(streetNames))] + " "
	fakeAddress += streetTypes[rand.Intn(len(streetTypes))] + ", "
	fakeAddress += cities[rand.Intn(len(cities))] + ", "
	fakeAddress += states[rand.Intn(len(states))] + " "
	fakeAddress += strconv.Itoa(rand.Intn(10000))
	return fakeAddress
}

var bookCategories = map[string]string{
	"Fiction":         "A type of narrative, typically in prose, which describes imaginary events and people.",
	"Non-Fiction":     "A type of narrative that deals with real events and people, often based on extensive research.",
	"Science Fiction": "A type of fiction that deals with imaginative and futuristic concepts, such as advanced science and technology, space exploration, time travel, parallel universes, and extraterrestrial life.",
	"Mystery":         "A genre of fiction in which a puzzle or a crime is presented for the reader or viewer to solve.",
	"Romance":         "A genre of fiction that deals with love and personal relationships.",
	"Thriller":        "A genre of fiction that keeps the reader or viewer on the edge of their seat with excitement and suspense.",
	"Horror":          "A genre of fiction that aims to create a feeling of fear, dread, and panic in the reader or viewer.",
	"Comedy":          "A genre of fiction that is intended to make its audience laugh.",
	"Drama":           "A genre of fiction that is intended to be serious and emotional, often involving intense character development and conflict.",
	"Adventure":       "A genre of fiction that typically features a hero on a quest or mission, often involving danger and exciting exploits.",
}

func GenerateFakeDate(start time.Time, end time.Time) time.Time {
	rand.Seed(time.Now().UnixNano())
	delta := end.Sub(start)
	randomDuration := time.Duration(rand.Int63n(int64(delta)))
	return start.Add(randomDuration)
}
