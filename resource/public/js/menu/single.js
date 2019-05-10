$(function () {
    const addComboCategory = function () {
        const lastID = parseInt($(this).attr('data-last-combo-category-id'))
        const nextID = lastID + 1
        const $list = $('#combo-category-list')
        $list.append(`
        <li class="combo-category-list-item">
            <div id="combo-category-${nextID}" class="combo-category">
                <input type="text" name="combo_category[]" hidden>
                <div class="form-group">
                    <input type="text" name="combo_category[${nextID}][name]" class="form-control" placeholder="Combo category name">
                </div>
                <div class="form-group">
                    <select name="combo_category[${nextID}][condition]" class="custom-select">
                        <option value="0" selected>Required</option>
                        <option value="1">Chose up to 1</option>
                    </select>
                </div>
                <div class="combo-list-container">
                    <ul class="combo-list">
                        <input type="text" name="combo_category[${nextID}][combo][]" hidden>
                        <li class="combo-list-item">
                            <div class="form-group">
                                <input type="text" name="combo_category[${nextID}][combo][0][name]" class="form-control" placeholder="Combo name">
                            </div>
                            <div class="form-group">
                                <span class="menu-price-currency">+ JPY</span><input type="text" name="combo_category[${nextID}][combo][0][price]" class="form-control combo-price-input" placeholder="Combo price">
                            </div>
                        </li>
                    </ul>
                    <div class="ritty-btn-border add-combo-btn" data-combo-category-id="${nextID}" data-last-combo-id="0">
                        <h1 class="ritty-btn-border-title add-combo-btn-title">+ Add combo</h1>
                    </div>
                </div>
            </div>
        </li>
        `)

        $('.add-combo-btn').click(addCombo)
        $(this).attr('data-last-combo-category-id', nextID)
    }

    const addCombo = function () {
        const categoryListID = $(this).data('combo-category-id')
        const lastID = parseInt($(this).attr('data-last-combo-id'))
        const nextID = lastID + 1
        const $comboList = $(`#combo-category-${categoryListID}`).find('.combo-list')
        $comboList.append(`
        <li class="combo-list-item">
            <input type="text" name="combo_category[${categoryListID}][combo][]" hidden>
            <div class="form-group">
                <input type="text" name="combo_category[${categoryListID}][combo][${nextID}][name]" class="form-control" placeholder="Combo name">
            </div>
            <div class="form-group">
                <span class="menu-price-currency">+ JPY</span><input type="text" name="combo_category[${categoryListID}][combo][${nextID}][price]" class="form-control combo-price-input" placeholder="Combo price">
            </div>
        </li>
        `)
        $(this).attr('data-last-combo-id', nextID)
    }

    const showMenuImage = function (e) {
        const file = e.target.files[0]
        const reader = new FileReader()
        const $preview = $('.menu-img')
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

    const deleteMenu = function (event) {
        event.preventDefault()
        $('#method-input').val('DELETE')
        $('#menu-form').submit()
    }

    $('#add-combo-category-btn').click(addComboCategory)
    $('.add-combo-btn').click(addCombo)
    $('form').on('change', 'input[type="file"]', showMenuImage)
    $('#delete-menu-btn').click(deleteMenu)
})