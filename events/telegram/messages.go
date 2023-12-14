package telegram

const msgHelp = `Я могу показывать вакансии сайта hh.ru!
/search - для поиска вакансии 
/settings - для установки фильтров`

const (
	msgHello = "<strong>Привет!</strong> 👾\n\n" + msgHelp
	// msgSettings   = "<strong>Начало установки фильтров!</strong> 🤓"
	msgSearch          = "<strong>Введите название профессии 🤓</strong>"
	msgCity            = "<strong>Введите название города 🤓</strong>"
	msgSalary          = "<strong>Введите желаемую зарплату 🤓</strong>"
	msgExperience      = "<strong>Выберет опыт работы 🤓</strong>"
	msgSettingsSuccess = "<strong>Фильтры заданы 🤓</strong>"

	msgUnknownCommand = "<strong>Неизвестная команда</strong> 🤔"
)
