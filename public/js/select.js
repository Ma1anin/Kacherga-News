jQuery(($) => {
    $('.select').on('click', '.select__head', function () {
        if ($(this).hasClass('open')) {
            $(this).removeClass('open');
            $(this).next().fadeOut();
        } else {
            $('.select__head').removeClass('open');
            $('.select__list').fadeOut();
            $(this).addClass('open');
            $(this).next().fadeIn();
        }
    });

    $('.select').on('click', '.select__item', function () {
        $('.select__head').removeClass('open');
        $(this).parent().fadeOut();
        $(this).parent().prev().text($(this).text());
        $(this).parent().prev().prev().val($(this).text());
    });

    $(document).click(function (e) {
        if (!$(e.target).closest('.select').length) {
            $('.select__head').removeClass('open');
            $('.select__list').fadeOut();
        }
    });
});

var pl = document.getElementById('player');
pl.currentTime = 43.0;

document.addEventListener('DOMContentLoaded', () => {
    let audio = document.querySelector('player');

    player.volume = 0.05;
}, false);


// const selector = document.getElementById("selector");
// const joker = document.getElementById("soska");

// joker.style.visibility = "hidden";

// function changeOption() {
//     const selectedOption = selector.options[selector.selectedIndex];
//     console.log(selectedOption.value);
//     if (selectedOption.value == "Лайк Лизы") {
//         joker.style.visibility = "visible";
//     } else {
//         joker.style.visibility = "hidden";
//     }
// }

// selector.addEventListener("change", changeOption);
