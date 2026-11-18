function arrangeRadialHoneycomb() {
    const viewport = document.getElementById('viewport');
    const container = document.getElementById('hex-grid');
    const circles = container.querySelectorAll('.circle');
    if (circles.length === 0) return;

    const screenW = window.innerWidth;
    let size = 140;
    let expandScale = 2;

    if (screenW <= 480) {
        size = 84;
        expandScale = 1.5;
    } else if (screenW <= 768) {
        size = 105;
        expandScale = 1.7;
    } else if (screenW <= 1024) {
        size = 120;
        expandScale = 1.85;
    }

    document.documentElement.style.setProperty('--circle-size', `${size}px`);
    document.documentElement.style.setProperty('--expand-scale', expandScale);

    const ratio = size / 140;
    const gapX = Math.round(-10 * ratio);
    const gapY = Math.round(10 * ratio);
    
    const paddingX = screenW <= 768 ? Math.max(200, Math.round(screenW * 0.5)) : 1200;
    const paddingY = screenW <= 768 ? Math.max(200, Math.round(window.innerHeight * 0.4)) : 500;

    const dx = size + gapX;
    const dy = size + gapY;

    const dirs = [
        { dx: dx,  dy: dy / 2 },
        { dx: 0,   dy: dy },
        { dx: -dx, dy: dy / 2 },
        { dx: -dx, dy: -dy / 2 },
        { dx: 0,   dy: -dy },
        { dx: dx,  dy: -dy / 2 }
    ];

    const positions = [];

    circles.forEach((circle, index) => {
        let x = 0;
        let y = 0;

        if (index > 0) {
            let r = 1;
            while (index >= 3 * r * (r + 1) + 1) {
                r++;
            }

            const startIndex = 3 * r * (r - 1) + 1;
            const posInRing = index - startIndex;

            x = 0;
            y = -r * dy;

            const side = Math.floor(posInRing / r);
            const steps = posInRing % r;

            for (let i = 0; i < side; i++) {
                x += r * dirs[i].dx;
                y += r * dirs[i].dy;
            }

            if (side < 6) {
                x += steps * dirs[side].dx;
                y += steps * dirs[side].dy;
            }
        }
        positions.push({ x, y, el: circle });
    });

    let maxDistX = 0;
    let maxDistY = 0;

    positions.forEach(pos => {
        if (Math.abs(pos.x) > maxDistX) maxDistX = Math.abs(pos.x);
        if (Math.abs(pos.y) > maxDistY) maxDistY = Math.abs(pos.y);
    });

    const offsetX = maxDistX + paddingX;
    const offsetY = maxDistY + paddingY;

    positions.forEach(pos => {
        pos.el.style.left = `${pos.x + offsetX}px`;
        pos.el.style.top = `${pos.y + offsetY}px`;
    });

    const gridWidth = (maxDistX * 2 + size) + (paddingX * 2);
    const gridHeight = (maxDistY * 2 + size) + (paddingY * 2);

    container.style.width = `${gridWidth}px`;
    container.style.height = `${gridHeight}px`;

    viewport.scrollLeft = (gridWidth - viewport.clientWidth) / 2;
    viewport.scrollTop = (gridHeight - viewport.clientHeight) / 2;
}

function enableDragToScroll() {
    const viewport = document.getElementById('viewport');
    const container = document.getElementById('hex-grid');

    let isMouseDown = false;
    let isDragging = false;
    let preventClick = false;

    let startX, startY, scrollLeft, scrollTop;
    let dragTimer;

    // --- настройка курсора ---
    const DRAG_DELAY = 300;      // мс, сколько нужно удерживать для активации скролла
    const DRAG_THRESHOLD = 5;    // px, на сколько нужно сдвинуть мышь для немедленной активации

    viewport.addEventListener('mousedown', (e) => {

        // Блокировка GrabToScroll при активация кнопки
        if (overlay.classList.contains('is-active')) return;

        isMouseDown = true;
        isDragging = false;
        preventClick = false;

        startX = e.pageX - viewport.offsetLeft;
        startY = e.pageY - viewport.offsetTop;
        scrollLeft = viewport.scrollLeft;
        scrollTop = viewport.scrollTop;

        // Если держим мышь DRAG_DELAY и не отпускаем, включаем драг
        dragTimer = setTimeout(() => {
            if (isMouseDown) {
                isDragging = true;
                viewport.classList.add('is-dragging');
                }
            }, DRAG_DELAY);
        });

        viewport.addEventListener('mouseleave', () => {
            isMouseDown = false;
            isDragging = false;
            viewport.classList.remove('is-dragging');
            clearTimeout(dragTimer);
        });

        viewport.addEventListener('mouseup', () => {
            isMouseDown = false;
            clearTimeout(dragTimer);

            if (isDragging) {
                preventClick = true;
                isDragging = false;
                viewport.classList.remove('is-dragging');
                
                // Снимаем блокировку клика через 50мс
                setTimeout(() => { preventClick = false; }, 50);
            }
        });

    // Двигаем мышью
    viewport.addEventListener('mousemove', (e) => {
        if (!isMouseDown) return;

        const x = e.pageX - viewport.offsetLeft;
        const y = e.pageY - viewport.offsetTop;
        const walkX = x - startX; 
        const walkY = y - startY; 

        if (!isDragging) {
            if (Math.abs(walkX) > DRAG_THRESHOLD || Math.abs(walkY) > DRAG_THRESHOLD) {
                isDragging = true;
                viewport.classList.add('is-dragging');
                clearTimeout(dragTimer);
            } else {
                return;
            }
        }

        e.preventDefault();
        viewport.scrollLeft = scrollLeft - walkX;
        viewport.scrollTop = scrollTop - walkY;
    });

    // Не даем перетаскивать картинки
    viewport.addEventListener('dragstart', (e) => {
        if (e.target.tagName === 'IMG') {
            e.preventDefault();
        }
    });

    // Вызов плавной передвижения и оверлей
    const overlay = document.getElementById('overlay');
    const detailsPanel = document.getElementById('details-panel');

    // Закрытие профиля при клике по темному фону
    overlay.addEventListener('click', () => {
        const currentExpanded = container.querySelector('.circle.is-expanded');
        if (currentExpanded) {
            currentExpanded.classList.remove('is-expanded');
        }
        overlay.classList.remove('is-active');
        detailsPanel.classList.remove('is-active');
    });

    container.addEventListener('click', (e) => {
        const button = e.target.closest('.circle');
        
        if (!button || preventClick) {
            e.preventDefault();
            e.stopPropagation();
            return;
        }

        if (button.classList.contains('is-expanded')) {
            button.classList.remove('is-expanded');
            overlay.classList.remove('is-active');
            return;
        }

        const currentExpanded = container.querySelector('.circle.is-expanded');
        if (currentExpanded) {
            currentExpanded.classList.remove('is-expanded');
        }

        const rawId = button.getAttribute('data-id');
        const artistId = parseInt(rawId, 10);

        console.log(`Успешный клик по артисту ID: ${artistId}`);
        
        const { id, name, creation, album, members, activity } = button.dataset;

        detailsPanel.innerHTML = `
            <div class="info-side info-left">
            <h2 class="days-one-regular">${name}</h2>
            <span class="badge badge-${activity} neucha-regular"><strong>Статус:</strong> ${activity}</span>
                <p class="neucha-regular"><strong>Год создания:</strong> ${creation}</p>
                <p class="neucha-regular"><strong>Первый альбом:</strong> ${album}</p>
            </div>

            <div class="center-spacer">
                <a href="/artist?id=${id}" class="action-btn days-one-regular">Подробнее</a>
            </div>

            <div class="info-side info-right">
                <p class="neucha-regular"><strong>Участники:</strong></p>
                <p class="neucha-regular">${members}</p>
            </div>
        `;

        button.classList.add('is-expanded');
        overlay.classList.add('is-active');

        const buttonCenterX = button.offsetLeft + (button.offsetWidth / 2);
        const buttonCenterY = button.offsetTop + (button.offsetHeight / 2);

        const targetScrollLeft = buttonCenterX - (viewport.clientWidth / 2);
        const targetScrollTop = buttonCenterY - (viewport.clientHeight / 2);

        customSmoothScroll(viewport, targetScrollLeft, targetScrollTop, 400);
    });

    // Рандомный цвета при наведение курсором
    let lastHue = -1;

    container.addEventListener('mouseover', (e) => {
        const button = e.target.closest('.circle');

        if (!button || button.contains(e.relatedTarget)) return;

        let hue;
        do {
            hue = Math.floor(Math.random() * 360);
        } while (Math.abs(hue - lastHue) < 50);
        lastHue = hue;

        button.style.setProperty('--hover-glow', `hsla(${hue}, 95%, 68%, 0.9)`);
    });
}

// Функция настраиваемого скролла
// element - что скроллим
// targetX, targetY - координаты куда едем
// duration - время анимации в миллисекундах
function customSmoothScroll(element, targetX, targetY, duration) {
    const startX = element.scrollLeft;
    const startY = element.scrollTop;
    const distanceX = targetX - startX;
    const distanceY = targetY - startY;
    let startTime = null;

    // Функция плавности EaseInOutQuad
    const easeInOut = (t, b, c, d) => {
        t /= d / 2;
        if (t < 1) return c / 2 * t * t + b;
        t--;
        return -c / 2 * (t * (t - 2) - 1) + b;
    };

    const animation = (currentTime) => {
        if (startTime === null) startTime = currentTime;
        const timeElapsed = currentTime - startTime;

        const nextX = easeInOut(timeElapsed, startX, distanceX, duration);
        const nextY = easeInOut(timeElapsed, startY, distanceY, duration);

        element.scrollLeft = nextX;
        element.scrollTop = nextY;

        if (timeElapsed < duration) {
            requestAnimationFrame(animation);
        } else {
            element.scrollLeft = targetX;
            element.scrollTop = targetY;
        }
    };

    requestAnimationFrame(animation);
}

function animateCirclesSequentially() {
    const circles = document.querySelectorAll('#hex-grid .circle');
    const STAGGER_DELAY = 70;

    circles.forEach((circle, index) => {
        setTimeout(() => {
            circle.classList.add('is-visible');
        }, index * STAGGER_DELAY);
    });
}

function initFloatingNotes() {
    const bgContainer = document.querySelector('.background');
    if (!bgContainer) return;

    // Берем размер высоты заднего фона
    const containerHeight = bgContainer.offsetHeight || window.innerHeight;
    bgContainer.style.setProperty('--travel-y', `-${containerHeight + 120}px`);

    // Настройка заднего фона
    const symbols = ['♪', '♫', '♩', '♬', '♭', '♯'];
    const NOTE_COUNT = 100;

    const fragment = document.createDocumentFragment();

    for (let i = 0; i < NOTE_COUNT; i++) {
        const note = document.createElement('span');
        note.classList.add('floating-note');
        note.textContent = symbols[Math.floor(Math.random() * symbols.length)];

        const startX = (Math.random() * 96 + 2).toFixed(1);
        const duration = (Math.random() * 18 + 20).toFixed(1);
        const delay = -(Math.random() * duration).toFixed(1);
        const driftX = Math.floor(Math.random() * 180 - 90);
        const rotStart = Math.floor(Math.random() * 40 - 20);
        const rotEnd = rotStart + Math.floor(Math.random() * 60 - 30);
        const size = Math.floor(Math.random() * 22 + 25);
        
        const maxOpacity = (Math.random() * 0.5 + 0.15).toFixed(2);
        const hue = Math.floor(Math.random() * 360);

        note.style.left = `${startX}%`;
        note.style.fontSize = `${size}px`;
        note.style.setProperty('--note-color', `hsl(${hue}, 90%, 70%)`);
        note.style.setProperty('--duration', `${duration}s`);
        note.style.setProperty('--delay', `${delay}s`);
        note.style.setProperty('--drift-x', `${driftX}px`);
        note.style.setProperty('--rot-start', `${rotStart}deg`);
        note.style.setProperty('--rot-end', `${rotEnd}deg`);
        note.style.setProperty('--max-opacity', maxOpacity);

        fragment.appendChild(note);
    }

    bgContainer.appendChild(fragment);
}

const detailsPanel = document.getElementById('details-panel');

// Не даем кликам и выделению текста внутри карточек закрывать overlay
detailsPanel.addEventListener('click', (e) => {
  // Если клик или выделение текста было внутри .info-side не закрываем окно
    if (e.target.closest('.info-side')) {
        e.stopPropagation();
    }
});

document.addEventListener('DOMContentLoaded', () => {
    arrangeRadialHoneycomb();
    enableDragToScroll();
    animateCirclesSequentially();
    initFloatingNotes();
});