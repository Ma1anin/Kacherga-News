// jQuery(($) => {
//     $('.select').on('click', '.select__head', function () {
//         if ($(this).hasClass('open')) {
//             $(this).removeClass('open');
//             $(this).next().fadeOut();
//         } else {
//             $('.select__head').removeClass('open');
//             $('.select__list').fadeOut();
//             $(this).addClass('open');
//             $(this).next().fadeIn();
//         }
//     });

//     $('.select').on('click', '.select__item', function () {
//         $('.select__head').removeClass('open');
//         $(this).parent().fadeOut();
//         $(this).parent().prev().text($(this).text());
//         $(this).parent().prev().prev().val($(this).text());
//     });

//     $(document).click(function (e) {
//         if (!$(e.target).closest('.select').length) {
//             $('.select__head').removeClass('open');
//             $('.select__list').fadeOut();
//         }
//     });
// });

document.addEventListener('DOMContentLoaded', () => {
    const selectHeads = document.querySelectorAll('.select__head');
    const selectLists = document.querySelectorAll('.select__list');

    selectHeads.forEach(head => {
        head.addEventListener('click', () => {
            const list = head.nextElementSibling;
            if (list.style.display === 'block') {
                list.style.display = 'none';
                head.classList.remove('open');
            } else {
                selectLists.forEach(list => list.style.display = 'none');
                selectHeads.forEach(head => head.classList.remove('open'));
                list.style.display = 'block';
                head.classList.add('open');
            }
        });
    });

    selectLists.forEach(list => {
        list.addEventListener('click', (e) => {
            const item = e.target;
            if (item.classList.contains('select__item')) {
                const head = item.closest('.select').querySelector('.select__head');
                head.textContent = item.textContent;
                list.style.display = 'none';
                head.classList.remove('open');
            }
        });
    });

    document.addEventListener('click', (e) => {
        if (!e.target.closest('.select')) {
            selectLists.forEach(list => list.style.display = 'none');
            selectHeads.forEach(head => head.classList.remove('open'));
        }
    });
});

var pl = document.getElementById('player');
pl.currentTime = 43.0;

document.addEventListener('DOMContentLoaded', () => {
    let audio = document.querySelector('player');

    player.volume = 0.05;
}, false);