function toggleShow(id, dis) {
    let elem = document.getElementById(id)
    console.log(elem.style.display)
    if (elem.style.display == "none") {
        elem.style.display = dis
    } else {
        elem.style.display = "none"
    }
}

function toggleShowMore(id) {
    let elem = document.getElementById(id)
    if (elem.innerText == "Show More") {
        elem.innerText = "Hide"
    } else {
        elem.innerText = "Show More"
    }
}
