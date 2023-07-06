function removeFromDb(name) {
    fetch(`/${name}`, { method: 'Delete' }).then((res) => {
        if (res.status == 200) {
            window.location.pathname = '/'
        }
    })
}
