package models

// Info
const (
	InfoTelega    = "Сәлеметсіз бе, бұл жұмыс/қызметкерлерді іздеуге арналған Telegram-бот.\n Здравствуйте, это Telegram-bot по поиску работы и сотрудников.\n Тілді таңдаңыз/Выберите язык"
	InfoInKazakh  = "Мәліметтер:\n Аты: %s \n Нөмірі: %s\n қызметі:%s\n Жасы: %d\n Жыныс: %s\n"
	InfoInRussian = "Сведение данных:\nИмя: %s\nНомер: %s\nДеятельность: %s\nВозраст: %d\nПол: %s\n"
)

// Russian
const (
	Russian                 = "russian"
	ChoseRussian            = "Вы выбрали русский язык"
	RussianName             = "Введите имя: "
	RussianHello            = "Здравствуйте, "
	RussianNumberButton     = "Для регистрации бота нам понадобится ваш номер телефона."
	RussianNumberInfo       = "Вы получите SMS-уведомление с кодом. Введите его, пожалуйста."
	RussianNumberRetrieved  = "Ваш код получен, спасибо"
	RussianSearch           = "%s спасибо за регистрацию! Вы ищите работу или сотрудника?"
	RussianWorkButton       = "ищу работу"
	RussianEmployeeButton   = "ишу сотрудника"
	RussianAgeButton        = "ввести возраст"
	RussianAccompanyingMess = "Сопроводительное письмо"
	RussianAgeInfo          = "Осталось совсем немного до окончания регистрации. Введите, пожалуйста, свой возраст."
	RussianGender           = "Укажите пол"
	RussianGenderMale       = "Мужчина"
	RussianGenderFemale     = "Женский"
)

// Kazakh
const (
	Kazakh                 = "kazakh"
	ChoseKazakh            = "сіз Қазақ тілің таңдадыңыз"
	KazakhName             = "Атыңызды енгізіңіз"
	KazakhHello            = "Сәлем, "
	KazakhNumberButton     = "Ботқа тіркелу үшін бізге телефон нөміріңіз қажет."
	KazakhNumberInfo       = "Сізге коды бар SMS хабарлама келеді. Оны енгізіңіз"
	KazakhNumberRetrieved  = "Сіздің кодыңыз алынды, рахмет"
	KazakhSearch           = "%s тіркелгеніңіз үшін рахмет! Сіз жұмыс немесе қызметкер іздеп жүрсіз бе?"
	KazakhWorkButton       = "жұмыс іздеу"
	KazakhEmployeeButton   = "қызметкер іздеу"
	KazakhAgeButton        = "жасын енгізіңіз"
	KazakhAccompanyingMess = "Ілеспе хат"
	KazakhAgeInfo          = "Тіркеу аяқталғанға дейін аз ғана уақыт қалды. Жасыңызды енгізіңіз."
	KazakhGender           = "Жынысын көрсетіңіз"
	KazakhGenderMale       = "Ер адам"
	KazakhGenderFemale     = "Әйел"
)

const ()

var Fields = []string{"Торговля", "Общепит", "Строительство"}

var Field = make(map[int][]string)

func init() {
	Field[1] = []string{"Продавец-консультант", "Менеджер по продажам", "Мерчандайзер", "Кассир"}

	Field[2] = []string{"Повар", "Бариста", "Официант", "Бармен"}

	Field[3] = []string{"Строитель", "Архитектор", "Архитектор", "Сварщик"}
}
