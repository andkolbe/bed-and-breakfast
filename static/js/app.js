function Prompt() {
    // can we simplify this to be more like success and error?
    let toast = function (c) {
        const { // c will be overridden by these values if not specified
            msg = '',
            icon = 'success',
            position = 'top-end',
        } = c;

        const Toast = Swal.mixin({
            toast: true,
            title: msg,
            position,
            icon,
            showConfirmButton: false,
            timer: 3000,
            timerProgressBar: true,
            didOpen: (toast) => {
                toast.addEventListener('mouseenter', Swal.stopTimer)
                toast.addEventListener('mouseleave', Swal.resumeTimer)
            }
        })

        Toast.fire({})
    }

    let success = function (c) {
        const {
            msg = '',
            title = '',
            footer = '',
        } = c

        Swal.fire({
            icon: 'success',
            title,
            text: msg,
            footer,
        })

    }

    let error = function (c) {
        const {
            msg = '',
            title = '',
            footer = '',
        } = c

        Swal.fire({
            icon: 'error',
            title,
            text: msg,
            footer,
        })

    }

    async function custom(c) {
        const {
            icon = '',
            msg = '',
            title = '',
            showConfirmButton = true,
        } = c;

        const { value: result } = await Swal.fire({
            icon,
            title,
            html: msg,
            backdrop: false,
            focusConfirm: false,
            showCancelButton: true,
            showConfirmButton,
            willOpen: () => {
                if (c.willOpen !== undefined) {
                    c.willOpen()
                }
            },
            didOpen: () => {
                if (c.didOpen !== undefined) {
                    c.didOpen()
                }
            }
        })
        // if a result comes back after someone clicks on the dialogue box
        if (result) {
            // and that result is not because someone clicked cancel
            if (result.dismiss !== Swal.DismissReason.cancel) {
                // and the result does not equal an empty string
                if (result.value !== "") {
                    // and the result is not undefined
                    if (c.callback !== undefined) {
                        c.callback(result)
                    }
                } else {
                    c.callback(false)
                }
            } else {
                c.callback(false)
            }
        }

    }

    return {
        toast,
        success,
        error,
        custom,
    }
}