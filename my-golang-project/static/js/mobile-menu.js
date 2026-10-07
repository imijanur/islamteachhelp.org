(function () {
	'use strict';

	var BREAKPOINT = 1024;
	var FALLBACK_MENU_ID = 'ekit-megamenu-header-menu';

	function forEach(list, fn) {
		for (var i = 0; i < list.length; i++) {
			fn(list[i], i);
		}
	}

	function initWrapper(wrap) {
		var panels = wrap.querySelectorAll('.elementskit-menu-offcanvas-elements');
		var togglers = wrap.querySelectorAll('.elementskit-menu-toggler');
		var hamburgers = wrap.querySelectorAll('.elementskit-menu-hamburger');
		var closes = wrap.querySelectorAll('.elementskit-menu-close');

		if (panels.length < 1 || togglers.length < 1) {
			return;
		}

		var menuId = FALLBACK_MENU_ID;
		forEach(panels, function (panel) {
			if (panel.id) {
				menuId = panel.id;
			}
		});

		forEach(hamburgers, function (hamburger) {
			hamburger.setAttribute('aria-expanded', 'false');
			hamburger.setAttribute('aria-controls', menuId);
		});
		forEach(closes, function (close) {
			if (!close.getAttribute('aria-label')) {
				close.setAttribute('aria-label', 'Close menu');
			}
		});

		function isOpen() {
			for (var i = 0; i < panels.length; i++) {
				if (panels[i].classList.contains('active')) {
					return true;
				}
			}
			return false;
		}

		function sync() {
			var open = isOpen();
			forEach(hamburgers, function (hamburger) {
				hamburger.setAttribute('aria-expanded', open ? 'true' : 'false');
			});
			document.body.style.overflow = open ? 'hidden' : '';
		}

		function setOpen(open) {
			forEach(panels, function (panel) {
				if (open) {
					panel.classList.add('active');
				} else {
					panel.classList.remove('active');
				}
			});
			sync();
		}

		forEach(togglers, function (toggler) {
			toggler.addEventListener('click', function (e) {
				e.preventDefault();
				setOpen(!isOpen());
			});
		});

		forEach(panels, function (panel) {
			forEach(panel.querySelectorAll('a'), function (link) {
				link.addEventListener('click', function () {
					setOpen(false);
				});
			});
		});

		document.addEventListener('keydown', function (e) {
			if ((e.key === 'Escape' || e.key === 'Esc') && isOpen()) {
				setOpen(false);
			}
		});

		window.addEventListener('resize', function () {
			if (window.innerWidth > BREAKPOINT && isOpen()) {
				setOpen(false);
			}
		});
	}

	function init() {
		forEach(document.querySelectorAll('.ekit-wid-con'), initWrapper);
	}

	if (document.readyState === 'loading') {
		document.addEventListener('DOMContentLoaded', init);
	} else {
		init();
	}
})();
