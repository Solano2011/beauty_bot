package telegram

import tele "gopkg.in/telebot.v3"

var (
	Menu = &tele.ReplyMarkup{}

	// Главное меню
	BtnMyBooking = Menu.Data("📅 Моя запись", "btn_my_booking")
	BtnContacts  = Menu.Data("📍 Контакты", "btn_contacts")

	// Навигация
	BtnBackToMain    = Menu.Data("◀️ Назад в меню", "btn_back_main")
	BtnCancelBooking = Menu.Data("❌ Отменить запись", "btn_cancel_booking")

	// Эндпоинты
	BtnService = Menu.Data("", "service")

	// Подтверждение замены записи
	BtnConfirmReplace = Menu.Data("✅ Да, отменить старую", "confirm_replace")
	BtnKeepOldBooking = Menu.Data("❌ Нет, оставить", "keep_old")

	// Админка
	BtnAdminRefresh  = Menu.Data("🔄 Обновить сводку", "admin_refresh")
	BtnAdminResetAll = Menu.Data("🗑 Сбросить все записи", "admin_reset_all")
)

func BuildMainMenu() *tele.ReplyMarkup {
	m := &tele.ReplyMarkup{}
	m.Inline(
		m.Row(BtnMyBooking),
		m.Row(BtnContacts),
	)
	return m
}

func BuildServicesMenu() *tele.ReplyMarkup {
	m := &tele.ReplyMarkup{}
	
	// Пока одна услуга (захардкожена)
	btnNails := m.Data("💅 Наращивание ногтей", "service", "Наращивание ногтей")
	
	m.Inline(
		m.Row(btnNails),
		m.Row(BtnBackToMain),
	)
	return m
}

func BuildContactsMenu() *tele.ReplyMarkup {
	m := &tele.ReplyMarkup{}
	btnMap := m.URL("🗺 Открыть на Яндекс.Картах", "https://yandex.ru/maps")
	m.Inline(
		m.Row(btnMap),
		m.Row(BtnBackToMain),
	)
	return m
}

func BuildReplaceConfirmMenu() *tele.ReplyMarkup {
	m := &tele.ReplyMarkup{}
	m.Inline(
		m.Row(BtnConfirmReplace),
		m.Row(BtnKeepOldBooking),
	)
	return m
}

func BuildAdminMenu() *tele.ReplyMarkup {
	m := &tele.ReplyMarkup{}
	m.Inline(
		m.Row(BtnAdminRefresh),
		m.Row(BtnAdminResetAll),
	)
	return m
}
