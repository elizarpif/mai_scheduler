package utils

import (
	"fmt"
	"testing"
)

func TestParseHtml(t *testing.T) {
	html := `<body class="cbp-spmenu-push">








	


<div id="cbp-spmenu-s2" class="cbp-spmenu cbp-spmenu-vertical cbp-spmenu-right"></div>
<a href="#" class="cbp-spmenu-area-outer"></a>
<div id="message" class="j-hidden"></div>


<div id="mainbody">





	<nav class="navbar navbar-static-top hidden-xs" id="navbar-top">
		<div class="container">
			<div id="navbar-quicknav" class="navbar-collapse collapse">

				<ul class="nav navbar-nav navbar-left">
	<li class="dropdown b-cbp-spmenu-item"> <a href="#" class="dropdown-toggle special" data-toggle="dropdown" role="button" aria-haspopup="true" aria-expanded="false">Школьникам <span class="caret"></span></a>
	<ul class="dropdown-menu">
		<li><a href="https://priem.mai.ru/?from=mai.ru">Приёмная комиссия </a></li>
		<li><a href="https://pre.mai.ru/">Обучение и мероприятия для школьников <sup style="font-weight:bold;color:orange;">NEW</sup></a></li>
		<li><a href="/education/napravleniya_podgotovki/">Направления подготовки</a></li>
		<li><a href="https://pre.mai.ru/admission/courses/?from=mairu">Подготовительные курсы</a></li>
		<li><a href="https://pre.mai.ru/events/public/?from=mairu">Дни открытых дверей</a></li>
		<li><a href="/education/price/">Стоимость обучения</a></li>
		<li><a href="/education/military/">Военная подготовка</a></li>
		<li><a href="https://pre.mai.ru/admission/fmshmai/?from=mairu">Физико-математическая школа</a> </li>
		<li><a href="/common/campus/">Кампус</a></li>
		<li><a href="https://preduniversariy.mai.ru/">Предуниверсарий МАИ</a></li>
		<li><a href="https://traektoria.mai.ru/">Детский технопарк</a></li>
	</ul>
 </li>
	<li class="dropdown b-cbp-spmenu-item"> <a href="#" class="dropdown-toggle" data-toggle="dropdown" role="button" aria-haspopup="true" aria-expanded="false">Студентам <span class="caret"></span></a>
	<ul class="dropdown-menu">
		<li><a href="/education/schedule/">Расписание занятий</a></li>
		<li><a href="http://lms.mai.ru">Электронная образовательная среда</a></li>
		<li><a href="/education/scholarship/">Стипендии</a></li>
		<li><a href="/life/welfare/sks/index.php">Оформление социальной карты студента</a></li>
		<li><a href="/education/international/">Академическая мобильность</a></li>
		<li><a href="/education/maisjtu/">Совместный институт МАИ-ШУЦТ</a></li>
		<li><a href="/education/courses.php">Курсы и дополнительное образование</a></li>
		<li><a href="/science/postgraduate/">Аспирантура</a></li>
		<li><a href="/common/campus/library/">Библиотека</a></li>
		<li><a href="/common/campus/">Кампус</a></li>
		<li><a href="http://job.mai.ru/">Отдел практик и трудоустройства</a></li>
		<li><a href="/cnp/">Центр начинающего предпринимательства</a></li>
		<li><a href="/life/join/index.php">Студенческие объединения</a></li>
		<li><a href="/science/nirs/index.php">Студенческая наука</a></li>
		<li><a href="/life/create/">Творческие коллективы и клубы</a></li>
	</ul>
 </li>
	<li class="dropdown b-cbp-spmenu-item"> <a href="#" class="dropdown-toggle" data-toggle="dropdown" role="button" aria-haspopup="true" aria-expanded="false">Сотрудникам <span class="caret"></span></a>
	<ul class="dropdown-menu">
		<li><a href="/unit/pers_dep/">Отдел кадров (кадровые документы)</a></li>
		<li><a href="/events/elections/">Выборы и конкурсы на замещение должностей</a></li>
		<li><a href="/other/effektivnyy-kontrakt/">Эффективный контракт</a></li>
		<li><a href="/life/associations/profcom_sotr/">Профсоюзная организация работников</a></li>
		<li><a href="/education/fpk/">Курсы и дополнительное образование</a></li>
		<li><a href="/science/services/">Технопарк МАИ</a></li>
		<li><a href="/publications/norm_docs/">Нормативные документы</a></li>
		<li><a href="/common/unit/reestr_smk.php">Реестр документации СМК</a></li>
	</ul>
 </li>
	<li class="dropdown b-cbp-spmenu-item"> <a href="#" class="dropdown-toggle" data-toggle="dropdown" role="button" aria-haspopup="true" aria-expanded="false">Выпускникам <span class="caret"></span></a>
	<ul class="dropdown-menu">
		<li><a href="http://www.clubmai.ru/">Клуб выпускников МАИ</a></li>
		<li><a href="http://job.mai.ru/">Отдел практик и трудоустройства</a></li>
		<li><a href="/science/postgraduate/">Аспирантура</a></li>
		<li><a href="/education/fpk/">Курсы и дополнительное образование</a></li>
		<li><a href="/education/eduacationglobal/">Программа «Глобальное образование»</a></li>
	</ul>
 </li>
	<li class="dropdown b-cbp-spmenu-item"> <a href="#" class="dropdown-toggle" data-toggle="dropdown" role="button" aria-haspopup="true" aria-expanded="false">Партнёрам <span class="caret"></span></a>
	<ul class="dropdown-menu">
		<li><a href="/science/services/">Услуги технопарка</a></li>
		<li><a href="/education/fpk/">Корпоративное обучение</a></li>
		<li><a href="http://conference-hall.mai.ru/">Аренда конференц-зала</a></li>
		<li><a href="/press/board/">Объявления</a></li>
		<li><a href="/press/pr/">Отдел по связям с общественностью</a></li>
	</ul>
 </li>
</ul>	

<ul class="nav navbar-nav navbar-right">


	<!--<li class="b-cbp-spmenu-item specproj text-bold"><a target="_blank"  data-clone="true" href="http://aeroweek.ru/?from=banner#registr" class="" style="">ASW</a></li>	-->
	<li class="b-cbp-spmenu-item specproj text-bold"><a class="" href="http://90.mai.ru?from=topbar">МАИ-90</a></li>
	<li class="b-cbp-spmenu-item specproj"><a class="" href="http://mai.ru/other/rector/?from=topbar#start">Интернет-приёмная</a></li>
	<!--<li class="separator">&nbsp;</li>-->

<!--data-bgimage-src="/bitrix/templates/mai15/img/content/icon-vzlet-2016.png" data-clone="true" href="http://launch.mai.ru/" class="b-icon b-icon-text" style="background-image: url('/bitrix/templates/mai15/img/content/icon-vzlet-2016.png');"-->


	
	
	<!--<li class="b-cbp-spmenu-item specproj"><a target="_blank" href="https://mai.ru/1september/?from=topbar" >#welcomeonboard</a></li>-->

	
	<li class="b-cbp-spmenu-item specproj"><a target="_blank"  data-clone="true" href="http://www.mai.ru/newleader?from=topbar" class="" style="">Школа управления</a></li>

	<li class="b-cbp-spmenu-item specproj"><a target="_blank"  data-clone="true" href="https://pay.mai.ru/" >PayMAI</a></li>
	<!--<li class="separator">&nbsp;</li>-->
	<li class="hidden-xs">
		<a title="English" href="http://en.mai.ru/" class="j-inline-block" style="width:40px;"><img style="margin-top:-2px;display:inline-block;" src="/bitrix/templates/mai15/img/icon-flag-gb.gif"></a>
	</li>
	<li class="hidden-xs">
		<a title="中文" href="http://cn.mai.ru/" class="j-inline-block" style="width:40px;"><img style="margin-top:-2px;display:inline-block;" src="/bitrix/templates/mai15/img/icon-flag-cn.gif"></a>
	</li>	
	<!--
	<li class="dropdown hidden-xs">
		<a class="lang-switcher" href="#" class="dropdown-toggle" data-toggle="dropdown" role="button" aria-haspopup="true" aria-expanded="false"><span class="j-nowrap j-inline-block"><img src="/bitrix/templates/mai15/img/icon-flag-ru.gif"><span class="caret"></span></span></a>
		<ul class="dropdown-menu">
			<li><a href="http://en.mai.ru/"><img src="/bitrix/templates/mai15/img/icon-flag-gb.gif"> English</a></li>
			<li><a href="http://intstudy.mai.ru/es"><img src="/bitrix/templates/mai15/img/icon-flag-spa.gif"> Espanol</a></li>
			<li><a href="http://intstudy.mai.ru/zh"><img src="/bitrix/templates/mai15/img/icon-flag-cn.gif"> Chineese</a></li>				
			<li><a href="http://en.mai.ru/kr/"><img src="/bitrix/templates/mai15/img/icon-flag-ko.gif"> Korean</a></li>
			<li><a href="http://intstudy.mai.ru/ae"><img src="/bitrix/templates/mai15/img/icon-flag-ae.gif"> Arabic</a></li>
			<li class="hidden"><a href="http://mai.ru/"><img src="/bitrix/templates/mai15/img/icon-flag-ru.gif"> Русский</a></li>
		</ul>
	</li>
	-->
	<li><a class="b-icon" title="Перейти на версию сайта для слабовидящих" href="/blind.html" itemprop="Copy"><span class="hidden-xs glyphicon glyphicon-eye-open"></span><span class="visible-xs">Для слабовидящих</span></a></li>

	
	<li id="user-profile-link" class="dropdown hidden-xs"><a class="dropdown-toggle b-icon bs-tooltip" href="#" data-placement="bottom" data-toggle="dropdown" role="button" aria-haspopup="true" aria-expanded="false" data-original-title="Елизавета Пивоварова (yapivov)"><span class="hidden-xs glyphicon glyphicon-user"></span><span class="visible-xs">yapivov</span><span class="caret"></span></a>
	
		<ul class="dropdown-menu">
			<!--<li><a href="/personal/profile/?backurl=%2Feducation%2Fschedule%2Findex.php%3Fdepartment%3D151%26course%3D1">yapivov</a></li>-->
			<li><a href="?logout=yes">Выйти</a></li>
		</ul>	
	
	</li>
	<!--
	<li id="user-logout-link"><a class="b-icon" href="?logout=yes" data-toggle="tooltip" data-placement="bottom" title="Выйти"><span class="hidden-xs glyphicon glyphicon-log-out"></span><span class="visible-xs">Выйти</span></a></li>-->
	<li data-pushmenu-display="false" class="visible-sm j-marg-left">
		<form action="/search/">
		<div class="navbar-search">
			<input type="text" name="q" class="form-control navbar-search-textinput" placeholder="Поиск"><div class="navbar-search-btn-wrap"><button type="submit" class="btn btn-default"><span class="glyphicon glyphicon-search"></span></button></div>
		</div>
		<input type="hidden" name="cx" value="000267500624155255505:qznp4y-wfpa" />
		<input type="hidden" name="cof" value="FORID:10" />
		<input type="hidden" name="ie" value="utf-8" />
		</form>
	</li>
</ul>
			</div>
		</div>
	</nav>

	<nav class="navbar navbar-static-top navbar-mai" role="navigation" id="navbar-mainmenu">
		<div class="container">
			<div class="visible-xs">
	
				<button type="button" class="navbar-toggle collapsed" id="pushmenu-toggler">
					<span class="sr-only">Toggle navigation</span>
					<span class="icon-bar"></span>
					<span class="icon-bar"></span>
					<span class="icon-bar"></span>
				</button>
	
				<a id="btn-contacts" class="navbar-xs-btn btn btn-default" href="/common/contacts/"><span class="glyphicon glyphicon-earphone"></span></a>		
				<button type="button" id="search-toggler" class="navbar-xs-btn btn btn-default" data-toggle="collapse" data-target="#searchform-xs" aria-controls="searchform-xs" aria-expanded="false" >
					 <span class="glyphicon glyphicon-search"></span>
				</button>
				
				<a class="visible-xs navbar-brand" href="https://mai.ru/"><img width="70" height="70" src="/bitrix/templates/mai15/img/logo-MAI90.svg" /></a>
			</div>
			<div id="navbar">


				

				<div class="hidden-xs" id="navbar-logo-sm"><a style="display:block;" href="https://mai.ru/"><img class="img-responsive" src="/bitrix/templates/mai15/img/logo-MAI90-border.png" /></a></div>



								
<div class="navbar-collapse collapse">
	<ul class="nav navbar-nav nav-justified">

						
		
			<!-- mouseover fix href="/common/"-->
			<!-- to remove dropdown remove next: data-toggle="dropdown" aria-haspopup="true" href="#"-->
			
			
			<li class="dropdown "><a href="/common/" class="dropdown-toggle" role="button" aria-expanded="false" ><span class="j-nowrap">МАИ</span> <span class="caret"></span></a>

				<ul class="dropdown-menu" role="menu">
										<li class=""><a href="/common/about/">Об университете</a></li>
										<li class=""><a href="/common/rector/">Приветствие ректора</a></li>
										<li class=""><a href="/common/leaders/">Ректорат</a></li>
										<li class=""><a href="/common/unit/index.php">Структура</a></li>
										<li class=""><a href="/common/program/">Программы развития</a>
				<ul>
										<li class=""><a href="/common/program/niu/">Национальный исследовательский университет</a></li>
										<li class=""><a href="/common/program/join/">Реорганизация МАИ и МАТИ</a></li>
										<li class=""><a href="/common/program/219/">Инновационная инфраструктура</a></li>
						</ul>
							<li class=""><a href="/common/achievments/">Достижения</a>
				<ul>
										<li class=""><a href="/common/achievments/ratings/">МАИ в рейтингах</a></li>
										<li class=""><a href="/common/achievments/index.php">Премии в области науки и техники</a></li>
										<li class=""><a href="/common/achievments/awardsedu/">Премии и награды в области образования</a></li>
						</ul>
							<li class=""><a href="/common/sotrudnichestvo/">Стратегическое партнёрство</a>
				<ul>
										<li class=""><a href="/common/sotrudnichestvo/index.php">Российские компании</a></li>
										<li class=""><a href="/common/sotrudnichestvo/scientific/">Российские образовательные и научные организации</a></li>
										<li class=""><a href="/common/sotrudnichestvo/foreing_company/">Зарубежные организации</a></li>
										<li class=""><a href="/common/sotrudnichestvo/membership/">Членство МАИ в ассоциациях и организациях</a></li>
										<li class=""><a href="/common/sotrudnichestvo/tp/">Технологические платформы</a></li>
										<li class=""><a href="/common/sotrudnichestvo/setevoe-vzaimodeystvie/">Сетевое взаимодействие</a></li>
										<li class=""><a href="/common/sotrudnichestvo/basic_chairs/">Базовые кафедры</a></li>
						</ul>
							<li class=""><a href="/common/campus/">Кампус</a>
				<ul>
										<li class=""><a href="/common/campus/dormitory.php">Общежития</a></li>
										<li class=""><a href="/common/campus/library/">Библиотека</a></li>
										<li class=""><a href="/common/campus/cafeteria/">Столовые, кафе</a></li>
										<li class=""><a href="/common/campus/sport/">Спортивная инфраструктура</a></li>
										<li class=""><a href="/common/campus/recreation/">Базы отдыха</a></li>
										<li class=""><a href="/common/campus/health/">Санаторий-профилакторий</a></li>
						</ul>
							<li class=""><a href="/common/history/">История</a>
				<ul>
										<li class=""><a href="/common/history/index.php">Краткая историческая справка</a></li>
										<li class=""><a href="/common/history/periods/">Годы и основные события в жизни МАИ</a></li>
										<li class=""><a href="/common/history/alumni/">Выпускники МАИ</a></li>
										<li class=""><a href="/common/history/years.php">Из истории высшего авиационного образования в России</a></li>
										<li class=""><a href="/common/history/honor_gallery.php">Галерея почёта</a></li>
										<li class=""><a href="/common/history/names">Названия МАИ</a></li>
										<li class=""><a href="/common/history/70-victory/">Ко Дню Победы</a></li>
										<li class=""><a href="/common/history/mai-85/">МАИ — 85</a></li>
										<li class=""><a href="/common/history/80/">МАИ — 80</a></li>
						</ul>
							<li class=""><a href="/common/brand/">Символика</a>
				<ul>
										<li class=""><a href="/common/brand/index.php">Эмблема и логотип МАИ</a></li>
										<li class=""><a href="/common/brand/linkexchange/">Баннеры сайта МАИ</a></li>
						</ul>
							<li class=""><a href="/common/documents/">Официальные документы</a>
				<ul>
										<li class=""><a href="/common/documents/index.php">Устав, лицензия, аккредитация</a></li>
										<li class=""><a href="/common/documents/reports/">Отчётность</a></li>
										<li class=""><a href="/common/documents/income/">Сведения о доходах</a></li>
										<li class=""><a href="/common/documents/norm_docs/">Нормативные документы</a></li>
						</ul>
							<li class=""><a href="/sveden/">Сведения об образовательной организации</a>
				<ul>
										<li class=""><a href="/sveden/common/">Основные сведения</a></li>
										<li class=""><a href="/sveden/struct/">Структура и органы управления образовательной организацией</a></li>
										<li class=""><a href="/sveden/document/">Документы</a></li>
										<li class=""><a href="/sveden/education/">Образование</a></li>
										<li class=""><a href="/sveden/eduStandarts/">Образовательные стандарты</a></li>
										<li class=""><a href="/sveden/employees">Руководство / педагогический (научно-педагогический) состав</a></li>
										<li class=""><a href="/sveden/objects/">Материально-техническое обеспечение и оснащённость образовательного процесса</a></li>
										<li class=""><a href="/sveden/grants/">Стипендии и иные виды материальной поддержки</a></li>
										<li class=""><a href="/sveden/paid_edu/">Платные образовательные услуги</a></li>
										<li class=""><a href="/sveden/budget/">Финансово-хозяйственная деятельность</a></li>
										<li class=""><a href="/sveden/vacant/">Вакантные места для приема (перевода)</a></li>
						</ul>
			</ul>
						
		
			<!-- mouseover fix href="/education/"-->
			<!-- to remove dropdown remove next: data-toggle="dropdown" aria-haspopup="true" href="#"-->
			
			
			<li class="dropdown item-selected"><a href="/education/" class="dropdown-toggle" role="button" aria-expanded="false" ><span class="j-nowrap">Образование</span> <span class="caret"></span></a>

				<ul class="dropdown-menu" role="menu">
										<li class=""><a href="/education/departments/">Институты, филиалы</a></li>
										<li class=""><a href="/education/napravleniya_podgotovki/">Направления подготовки</a></li>
										<li class=""><a href="/education/price/">Стоимость обучения</a>
				<ul>
										<li class=""><a href="/education/price/firstcourse/">Стоимость обучения для студентов 1 курса</a></li>
										<li class=""><a href="/education/price/sophomore/">Стоимость обучения для студентов 2-го и последующих курсов</a></li>
										<li class=""><a href="/education/price/rktmai.php">Стоимость обучения для студентов филиала «РКТ» МАИ</a></li>
						</ul>
							<li class=""><a href="/education/fpk/">Повышение квалификации и переподготовка</a></li>
										<li class=""><a href="/education/schoolclubs/">Кружки, курсы и мероприятия для школьников</a>
				<ul>
										<li class=""><a href="/education/schoolclubs/index.php">Кружки, курсы и мероприятия для школьников</a></li>
										<li class=""><a href="/education/schoolclubs/opendays/">Дни открытых дверей</a></li>
						</ul>
							<li class=""><a href="/education/scholarship/">Стипендии</a></li>
										<li class=""><a href="/education/international/">Академическая мобильность</a>
				<ul>
										<li class=""><a href="http://mai.ru/education/international/">Академическая мобильность</a></li>
										<li class=""><a href="/education/international/directinternships/">Прямые стажировки от МАИ</a></li>
										<li class=""><a href="/education/international/faq.php">F.A.Q.</a></li>
						</ul>
							<li class=""><a href="/education/maisjtu/">Совместный институт МАИ-ШУЦТ</a>
				<ul>
										<li class=""><a href="/education/maisjtu/master/">Магистратура МАИ-ШУЦТ</a></li>
										<li class=""><a href="/education/maisjtu/bachelor/">Бакалавриат МАИ-ШУЦТ</a></li>
										<li class=""><a href="/education/maisjtu/summer_school/">Летняя школа в ШУЦТ</a></li>
						</ul>
							<li class=""><a href="https://mai.ru/itmai/">IT-центр</a></li>
										<li class="item-selected"><a href="/education/schedule/">Расписание занятий</a></li>
						</ul>
						
		
			<!-- mouseover fix href="/science/"-->
			<!-- to remove dropdown remove next: data-toggle="dropdown" aria-haspopup="true" href="#"-->
			
			
			<li class="dropdown "><a href="/science/" class="dropdown-toggle" role="button" aria-expanded="false" ><span class="j-nowrap">Наука</span> <span class="caret"></span></a>

				<ul class="dropdown-menu" role="menu">
										<li class=""><a href="/science/postgraduate/">Аспирантура</a></li>
										<li class=""><a href="/science/dissovet/">Диссертационные советы</a></li>
										<li class=""><a href="/science/results/">Результаты научной деятельности</a>
				<ul>
										<li class=""><a href="/science/results/dev/">Разработки</a></li>
										<li class=""><a href="/science/results/patents/">Патенты и свидетельства</a></li>
						</ul>
							<li class=""><a href="/science/projects/">Проекты</a>
				<ul>
										<li class=""><a href="/science/projects/gov_projects.php">Проекты в рамках Постановлений Правительства РФ № 218 и № 220</a></li>
										<li class=""><a href="/science/projects/grants">Гранты Президента Российской Федерации</a></li>
										<li class=""><a href="/science/projects/schools/">Научные школы</a></li>
										<li class=""><a href="/science/projects/fcp/">Федеральные целевые программы</a></li>
						</ul>
							<li class=""><a href="/science/services/">Научно-технический парк</a></li>
										<li class=""><a href="/science/bi/">Бизнес-инкубатор</a>
				<ul>
										<li class=""><a href="http://www.mai.ru/science/bi/mip/">Малые инновационные предприятия</a></li>
										<li class=""><a href="http://www.mai.ru/science/bi/pu/">Программа «УМНИК»</a></li>
						</ul>
							<li class=""><a href="/science/cpg/">Конкурсы, программы, гранты</a></li>
										<li class=""><a href="/science/publish/">Научные издания</a></li>
										<li class=""><a href="/science/nirs/">Студенческая наука</a></li>
										<li class=""><a href="/science/conf.php">Научные мероприятия МАИ</a></li>
										<li class=""><a href="/science/date">Внешние научные мероприятия</a></li>
										<li class=""><a href="/science/author/">Информация для авторов</a>
				<ul>
										<li class=""><a href="/science/author/index.php">Информация для авторов публикаций</a></li>
										<li class=""><a href="/science/author/contract/">Публикации и эффективный контракт</a></li>
										<li class=""><a href="/science/author/article/">Рекомендации по подготовке и оформлению статей</a></li>
										<li class=""><a href="/science/author/base/">Рекомендации по работе с информационными базами</a></li>
										<li class=""><a href="/science/author/list/">Журналы</a></li>
										<li class=""><a href="/science/author/board">Объявления</a></li>
										<li class=""><a href="/science/author/contact/">Контактная информация</a></li>
						</ul>
			</ul>
						
		
			<!-- mouseover fix href="/life/"-->
			<!-- to remove dropdown remove next: data-toggle="dropdown" aria-haspopup="true" href="#"-->
			
			
			<li class="dropdown "><a href="/life/" class="dropdown-toggle" role="button" aria-expanded="false" ><span class="j-nowrap">Жизнь</span> <span class="caret"></span></a>

				<ul class="dropdown-menu" role="menu">
										<li class=""><a href="/life/events/index.php">События в жизни МАИ</a></li>
										<li class=""><a href="/life/join/index.php">Студенческие объединения</a></li>
										<li class=""><a href="/life/sport/sections.php">Спортивные секции</a></li>
										<li class=""><a href="/life/create/">Творческие коллективы и клубы</a>
				<ul>
										<li class=""><a href="/life/create/dkit/index.php">Дворец культуры и техники</a></li>
										<li class=""><a href="/life/create/dkit/kollektivy-dkit.php">Творческие коллективы</a></li>
						</ul>
							<li class=""><a href="/life/welfare/">Материальная поддержка</a>
				<ul>
										<li class=""><a href="/life/welfare/sks/index.php">Оформление социальной карты студента</a></li>
						</ul>
							<li class=""><a href="/life/associations/">Объединения сотрудников и выпускников</a></li>
						</ul>
						
		
			<!-- mouseover fix href="/press/"-->
			<!-- to remove dropdown remove next: data-toggle="dropdown" aria-haspopup="true" href="#"-->
			
			
			<li class="dropdown "><a href="/press/" class="dropdown-toggle" role="button" aria-expanded="false" ><span class="j-nowrap">Пресс-центр</span> <span class="caret"></span></a>

				<ul class="dropdown-menu" role="menu">
										<li class=""><a href="/press/news/">Новости</a></li>
										<li class=""><a href="/press/events/">Мероприятия</a></li>
										<li class=""><a href="/press/board/">Объявления</a></li>
										<li class=""><a href="/press/dates/">Юбилеи</a></li>
										<li class=""><a href="/press/iam/">Маёвские истории успеха</a></li>
										<li class=""><a href="/press/media/">Медиабанк</a>
				<ul>
										<li class=""><a href="/press/media/photo/">Фото</a></li>
										<li class=""><a href="/press/media/presentations/">Презентации</a></li>
										<li class=""><a href="/press/media/video/">Видео</a></li>
										<li class=""><a href="/press/media/audio/">Аудио</a></li>
						</ul>
							<li class=""><a href="/press/print/">Корпоративные СМИ</a>
				<ul>
										<li class=""><a href="/press/print/cloud/">Облако</a></li>
										<li class=""><a href="/press/print/propeller/">Пропеллер</a></li>
						</ul>
							<li class=""><a href="/press/pr/">Отдел по связям с общественностью</a></li>
										<li class=""><a href="/press/smi/">Аккредитация СМИ</a></li>
						</ul>
						
		
			<!-- mouseover fix href="/common/contacts/"-->
			<!-- to remove dropdown remove next: data-toggle="dropdown" aria-haspopup="true" href="#"-->
			
			
			<li class="dropdown "><a href="/common/contacts/" class="dropdown-toggle" role="button" aria-expanded="false" ><span class="j-nowrap">Контакты</span> <span class="caret"></span></a>

				<ul class="dropdown-menu" role="menu">
										<li class=""><a href="/common/contacts/index.php">Контакты, схема проезда</a></li>
										<li class=""><a href="/common/contacts/bank.php">Банковские реквизиты</a></li>
						</ul></li>

		<li data-pushmenu-display="false" class="hidden-sm">
			<form action="/search/">
			<div class="navbar-search">

				<input type="text" name="q" class="form-control navbar-search-textinput" placeholder="Поиск"><div class="navbar-search-btn-wrap"><button type="submit" class="btn btn-default"><span class="glyphicon glyphicon-search"></span></button></div>
			</div>
			<input type="hidden" name="cx" value="000267500624155255505:qznp4y-wfpa" />
			<input type="hidden" name="cof" value="FORID:10" />
			<input type="hidden" name="ie" value="utf-8" />
			</form>
		</li>
		


</ul>
</div>

			</div>
		</div>
	</nav><!-- /block navbar -->
	
	
	


	
	
	<div class="j-bg-white collapse" id="searchform-xs" aria-expanded="false">
		<div class="container">
			<div class="row j-padd2-md">
				<div class="col-md-12">
				<form action="/search/">
					<input type="hidden" name="cx" value="000267500624155255505:qznp4y-wfpa" />
					<input type="hidden" name="cof" value="FORID:10" />
					<input type="hidden" name="ie" value="utf-8" />
					<div id="mobi-search"><input id="searchform-xs-input" type="text" name="q" class="form-control" placeholder="Поиск"><div id="mobi-search-btn-wrap"><button type="submit" class="btn btn-default">Найти</button></div></div>
				</form>
				</div>
			</div>
		</div>
	</div>

	
	
	
	
	
	
	
	
	
<style type="text/css">

#dod3
{
	border : none;
	background : #fff url(/bitrix/templates/mai15/components/mai/simple/slides/img/dod-3/main.png) no-repeat;
	min-height : 350px;
	padding-bottom : 30px;
	margin-bottom : 40px;
}
@media screen and (min-width: 1600px)
{

	#dod-3	{	background-position: 85% center;	}	
}
@media screen and (min-width: 1200px) and (max-width: 1599px)
{

	#dod3	{	background-position: 110% center;	}
}
@media screen and (min-width: 992px) and (max-width: 1199px)
{
	#dod3	{	background-position: 330px center;	}
}
@media screen and (min-width: 768px) and (max-width: 991px)
{
	#dod3 	{	background-position: 340px center;	}
}	
@media screen and (max-width: 767px)
{
	#dod3  {background : #fff;  margin-bottom : 0; min-height : 200px;}
} 
.b-textbody img[src$="imgonline_com_ua_Resize_mmcJtLE1LVA.jpg"]
{
	display : none !important;
}
</style>


			
	
	
	
	<div class="b-header" id="">

		<div class="container">
			<div class="row j-marg-top"><div class="col-md-12"><ol class="breadcrumb"><li><a href="/">Главная</a></li><li><a href="/education/">Образование</a></li><li><span>Расписание</span></li></ol></div></div>			<div class="row">
				<div class="col-md-12">
				
				<h3>Расписание</h3>
								</div>
			</div>
		</div>
	</div>
		
	
	
	
	
	
	




	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	<div id="b-submenu-tablet-btn-container" class="j-hidden">
		<div class="container">
			<div class="row">
				<div class="col-md-12"><a href="#" class="btn btn-default"><span class="glyphicon glyphicon-chevron-right j-color-white"></span></a></div>
			</div>
		</div>
	</div>	
	
	<div class="j-bg-mobi ">
		<div class="container">
			<div class="row text-lh-sm ">
				
				<div id="b-submenu-tablet" class="j-hidden col-sm-3 j-padd2-md"></div>
			
				<div style="margin-bottom:5vmax;" class="b-textbody j-marg-top		col-md-12">				
				
				<div id="schedule-content">


	<!--
	<div class="visible-xs j-marg-bottom">
		<span class="schedule-update">обновлено в 15:17</span><span class="schedule-today-date">8 сентября, чётная неделя</span></span>
	</div>
	-->

	<div>


		<form role="tabpanel" id="student">
		

			<div class="row hidden-xs">
				<label class="col-sm-4 col-md-2 control-label j-padd-top-xs" for="selectbasic">Факультет</label>


				<label class="col-sm-4 col-md-2 control-label j-padd-top-xs" for="course">Курс</label>
		  
			</div>
			<div class="row j-marg-bottom">
				<div class="col-sm-4 col-md-2">
				<label class="visible-xs j-padd-top-xs" for="selectbasic">Факультет</label>
				<select id="department" name="department" class="form-control">
				  <option value="0">Все факультеты</option>
				  				  <option value="150" >Факультет №1</option>
				  				  <option value="153" >Институт №2</option>
				  				  <option value="157" >Институт №3</option>
				  				  <option value="149" >Институт №4</option>
				  				  <option value="155" >Институт №5</option>
				  				  <option value="160" >Институт №6</option>
				  				  <option value="154" >Институт №7</option>
				  				  <option value="151" selected="selected">Институт №8</option>
				  				  <option value="152" >Институт №9</option>
				  				  <option value="146" >Институт №10</option>
				  				  <option value="165" >Институт №11</option>
				  				  <option value="164" >Институт №12</option>
				  				</select>
				</div>

				<div class="col-sm-4 col-md-2">
					<label class="visible-xs j-padd-top-xs" for="course">Курс</label>
					<select id="course" name="course" class="form-control">
					  <option value="0">Все курсы</option>
					  						<option value="1" selected="selected">1</option>
					  						<option value="2" >2</option>
					  						<option value="3" >3</option>
					  						<option value="4" >4</option>
					  						<option value="5" >5</option>
					  						<option value="6" >6</option>
					  					</select>
				</div>
				<div class="col-sm-4 col-md-2">
					<div class="visible-xs j-padd-top-xs" for="course">&nbsp;</div>
					<button type="submit" class="btn btn-primary">Отобразить</button>
				</div>				  
			</div>
				
		
		</form>
			
	</div>		































<!--
<a class="btn" href="#1">Курс 1</a>
<br><br>
-->


<div class="">

	<h2 class="result-header"></h2>

	
</div>



	<div class="sc-container">
		<h5 class="sc-container-header sc-gray" >1 курс</h5>
		<div class="sc-table sc-table-groups" id="1">
		
			<div class="sc-table-row">
				<a class="sc-table-col" href="#fac-1-Институт-№8" role="button" data-toggle="collapse" aria-expanded="false" aria-controls="fac-1-Институт-№8">Институт №8</a>
				<div class="sc-table-col">
					<div class="sc-table-col-body " id="fac-1-Институт-№8">
											<div class="sc-groups">
							<div class="sc-program">Аспирантура</div>
																<a class="sc-group-item" href="detail.php?group=М8О-101А-19">М8О-101А-19</a>

																<a class="sc-group-item" href="detail.php?group=М8О-102А-19">М8О-102А-19</a>

																<a class="sc-group-item" href="detail.php?group=М8О-104А-19">М8О-104А-19</a>

																<a class="sc-group-item" href="detail.php?group=М8О-105А-19">М8О-105А-19</a>

																<a class="sc-group-item" href="detail.php?group=М8О-106А-19">М8О-106А-19</a>

																<a class="sc-group-item" href="detail.php?group=М8О-108А-19">М8О-108А-19</a>

																<a class="sc-group-item" href="detail.php?group=М8О-111А-19">М8О-111А-19</a>

													
						</div>

											<div class="sc-groups">
							<div class="sc-program">Бакалавриат</div>
																<a class="sc-group-item" href="detail.php?group=М8О-101Б-19">М8О-101Б-19</a>

																<a class="sc-group-item" href="detail.php?group=М8О-102Б-19">М8О-102Б-19</a>

																<a class="sc-group-item" href="detail.php?group=М8О-103Б-19">М8О-103Б-19</a>

																<a class="sc-group-item" href="detail.php?group=М8О-104Б-19">М8О-104Б-19</a>

																<a class="sc-group-item" href="detail.php?group=М8О-105Б-19">М8О-105Б-19</a>

																<a class="sc-group-item" href="detail.php?group=М8О-106Б-19">М8О-106Б-19</a>

																<a class="sc-group-item" href="detail.php?group=М8О-107Б-19">М8О-107Б-19</a>

																<a class="sc-group-item" href="detail.php?group=М8О-108Б-19">М8О-108Б-19</a>

																<a class="sc-group-item" href="detail.php?group=М8О-109Б-19">М8О-109Б-19</a>

																<a class="sc-group-item" href="detail.php?group=М8О-110Б-19">М8О-110Б-19</a>

																<a class="sc-group-item" href="detail.php?group=М8О-111Б-19">М8О-111Б-19</a>

																<a class="sc-group-item" href="detail.php?group=М8О-112Б-19">М8О-112Б-19</a>

													
						</div>

											<div class="sc-groups">
							<div class="sc-program">Магистратура</div>
																<a class="sc-group-item" href="detail.php?group=М8В-101Мк-19">М8В-101Мк-19</a>

																<a class="sc-group-item" href="detail.php?group=М8О-101М-19">М8О-101М-19</a>

																<a class="sc-group-item" href="detail.php?group=М8О-102М-19">М8О-102М-19</a>

																<a class="sc-group-item" href="detail.php?group=М8О-103М-19">М8О-103М-19</a>

																<a class="sc-group-item" href="detail.php?group=М8О-104М-19">М8О-104М-19</a>

																<a class="sc-group-item" href="detail.php?group=М8О-105М-19">М8О-105М-19</a>

																<a class="sc-group-item" href="detail.php?group=М8О-106М-19">М8О-106М-19</a>

																<a class="sc-group-item" href="detail.php?group=М8О-107М-19">М8О-107М-19</a>

																<a class="sc-group-item" href="detail.php?group=М8О-108М-19">М8О-108М-19</a>

																<a class="sc-group-item" href="detail.php?group=М8О-110М-19">М8О-110М-19</a>

																<a class="sc-group-item" href="detail.php?group=М8О-111М-19">М8О-111М-19</a>

																<a class="sc-group-item" href="detail.php?group=М8О-113М-19">М8О-113М-19</a>

													
						</div>

										</div>
				</div>
			</div>

				</div>
	</div>
13:40:45<br></div>
				</div>
				
								
			</div>
		</div>
	</div>

	



	

	<!-- block footer -->
	<div id="footer" class="j-bg-lightblue j-bg-mobi">
		<div class="container">

			<div class="row ">

								

	<div class="hidden-xs col-sm-4 col-md-2 j-marg-bottom">
		<h5><a href="/common/">МАИ</a></h5>
			<ul class="j-vlist j-vlist-by-border text-sm">

				<li><a href="/common/about/">Об университете</a></li>

				<li><a href="/common/rector/">Приветствие ректора</a></li>

				<li><a href="/common/leaders/">Ректорат</a></li>

				<li><a href="/common/unit/index.php">Структура</a></li>

				<li><a href="/common/program/">Программы развития</a></li>

				<li><a href="/common/achievments/">Достижения</a></li>

				<li><a href="/common/sotrudnichestvo/">Стратегическое партнёрство</a></li>

				<li><a href="/common/campus/">Кампус</a></li>

				<li><a href="/common/history/">История</a></li>

				<li><a href="/common/brand/">Символика</a></li>

				<li><a href="/common/documents/">Официальные документы</a></li>

				<li><a href="/sveden/">Сведения об образовательной организации</a></li>
			</ul>
	</div>

	<div class="hidden-xs col-sm-4 col-md-2 j-marg-bottom">
		<h5><a href="/education/">Образование</a></h5>
			<ul class="j-vlist j-vlist-by-border text-sm">

				<li><a href="/education/departments/">Институты, филиалы</a></li>

				<li><a href="/education/napravleniya_podgotovki/">Направления подготовки</a></li>

				<li><a href="/education/price/">Стоимость обучения</a></li>

				<li><a href="/education/fpk/">Повышение квалификации и переподготовка</a></li>

				<li><a href="/education/schoolclubs/">Кружки, курсы и мероприятия для школьников</a></li>

				<li><a href="/education/scholarship/">Стипендии</a></li>

				<li><a href="/education/international/">Академическая мобильность</a></li>

				<li><a href="/education/maisjtu/">Совместный институт МАИ-ШУЦТ</a></li>

				<li><a href="https://mai.ru/itmai/">IT-центр</a></li>

				<li><a href="/education/schedule/">Расписание занятий</a></li>
			</ul>
	</div>

	<div class="hidden-xs col-sm-4 col-md-2 j-marg-bottom">
		<h5><a href="/science/">Наука</a></h5>
			<ul class="j-vlist j-vlist-by-border text-sm">

				<li><a href="/science/postgraduate/">Аспирантура</a></li>

				<li><a href="/science/dissovet/">Диссертационные советы</a></li>

				<li><a href="/science/results/">Результаты научной деятельности</a></li>

				<li><a href="/science/projects/">Проекты</a></li>

				<li><a href="/science/services/">Научно-технический парк</a></li>

				<li><a href="/science/bi/">Бизнес-инкубатор</a></li>

				<li><a href="/science/cpg/">Конкурсы, программы, гранты</a></li>

				<li><a href="/science/publish/">Научные издания</a></li>

				<li><a href="/science/nirs/">Студенческая наука</a></li>

				<li><a href="/science/conf.php">Научные мероприятия МАИ</a></li>

				<li><a href="/science/date">Внешние научные мероприятия</a></li>

				<li><a href="/science/author/">Информация для авторов</a></li>
			</ul>
	</div>

<div class="clearfix visible-sm"></div>	<div class="hidden-xs col-sm-4 col-md-2 j-marg-bottom">
		<h5><a href="/life/">Жизнь</a></h5>
			<ul class="j-vlist j-vlist-by-border text-sm">

				<li><a href="/life/events/index.php">События в жизни МАИ</a></li>

				<li><a href="/life/join/index.php">Студенческие объединения</a></li>

				<li><a href="/life/sport/sections.php">Спортивные секции</a></li>

				<li><a href="/life/create/">Творческие коллективы и клубы</a></li>

				<li><a href="/life/welfare/">Материальная поддержка</a></li>

				<li><a href="/life/associations/">Объединения сотрудников и выпускников</a></li>
			</ul>
	</div>

	<div class="hidden-xs col-sm-4 col-md-2 j-marg-bottom">
		<h5><a href="/press/">Пресс-центр</a></h5>
			<ul class="j-vlist j-vlist-by-border text-sm">

				<li><a href="/press/news/">Новости</a></li>

				<li><a href="/press/events/">Мероприятия</a></li>

				<li><a href="/press/board/">Объявления</a></li>

				<li><a href="/press/dates/">Юбилеи</a></li>

				<li><a href="/press/iam/">Маёвские истории успеха</a></li>

				<li><a href="/press/media/">Медиабанк</a></li>

				<li><a href="/press/print/">Корпоративные СМИ</a></li>

				<li><a href="/press/pr/">Отдел по связям с общественностью</a></li>

				<li><a href="/press/smi/">Аккредитация СМИ</a></li>
			</ul>
	</div>

	<div class="hidden-xs col-sm-4 col-md-2 j-marg-bottom">
		<h5><a href="/common/contacts/">Контакты</a></h5>
			<ul class="j-vlist j-vlist-by-border text-sm">

				<li><a href="/common/contacts/index.php">Контакты, схема проезда</a></li>

				<li><a href="/common/contacts/bank.php">Банковские реквизиты</a></li>
			</ul>
	</div>
				
				<div class="col-sm-4 col-md-2">


				
					<div class="row">
						<div class="col-xs-12 col-sm-12">
							<div class="text-md j-padd-top-sm">
								<ul class="hidden-xs j-vlist j-vlist-by-space">
									<li><a href="/common/contacts/">Контакты</a></li>
									<li><a href="/site/">О портале</a></li>
									<li><a href="/search/map.php">Карта сайта</a></li>
									<li><a href="/common/documents/income/">Сведения о доходах</a></li>
									<li><a href="/other/protivodeystvie-terrorizmu/">Противодействие терроризму</a></li>
									<li><a href="/other/protivodeystvie-korruptsii/index.php">Противодействие коррупции</a></li>
									
									<li class="j-hidden"><a id="btn-goto-mobile" href="#">Мобильная версия</a></li>
								</ul>
								<ul class="visible-xs j-vlist j-vlist-by-space">
									<li><h5><a href="/common/contacts/">Контакты</a></h5></li>
									<li><h5><a href="/site/">О портале</a></h5></li>
									<li><h5><a href="/search/map.php">Карта сайта</a></h5></li>
									<li><h5><a href="/common/documents/income/">Сведения о доходах</a></h5></li>
									<li><h5><a href="/other/protivodeystvie-terrorizmu/">Противодействие терроризму</a></h5></li>	
									<li><h5><a href="/other/protivodeystvie-korruptsii/index.php">Противодействие коррупции</a></h5></li>									
								</ul>								
							</div>						
						</div>
					</div>

				</div>
			</div>

			<div class="row">
				<div class="col-md-12 j-padd-top-lg text-xs  text-muted">© 1997—2020 МАИ</div>
			</div>

		</div>
	</div>
	<!-- /block footer -->

	
</div>



</body>`
	_, err := ParseHtml(html)
	if err !=nil{
		fmt.Println(err)
	} else{
		fmt.Print("success")
		//findR := res.FindAllSt("sc-groups")
		//for _, r := range findR{
		//	r.Print()
		//}

	}
}
