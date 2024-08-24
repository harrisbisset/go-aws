function toggleShow(id) {
    let elem = document.getElementById(id, dis)
    console.log(elem.style.display)
    if (elem.style.visibility == "none") {
        elem.style.visibility = dis
    } else {
        elem.style.visibility = "none"
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
