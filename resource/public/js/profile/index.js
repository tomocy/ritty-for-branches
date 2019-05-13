$(function () {
    const showMenuImage = function (e) {
        const file = e.target.files[0]
        const reader = new FileReader()
        const $preview = $('.img')
        const $btnTitle = $preview.find('.ritty-btn-border-title')

        if (file.type.indexOf('image') < 0) {
            return false
        }

        reader.onload = (function (file) {
            return function (e) {
                $preview.css('backgroundImage', `url('${e.target.result}')`)
            }
        })(file)

        reader.readAsDataURL(file)

        $preview.removeClass('d-none').addClass('d-block')
        $btnTitle.removeClass('d-block').addClass('d-none')
    }

    $('form').on('change', 'input[type="file"]', showMenuImage)
})