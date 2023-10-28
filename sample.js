// ------------- const ---------------

const server = "http://localhost:7777";


// ------------- helper ------------

async function post(target, body) {
    fetch(target, {
        method: 'POST',
        credentials: 'include',
        // header
        headers: {
            'Content-Type': 'application/json',
        },
        // json body
        body: JSON.stringify(body)
    }
    ).then(response => {
        if (response.ok) {
            response.json().then(json => console.log(json))
            return
        }
        throw new Error('Network response was not ok.');
    }).catch(err => console.log(err));
}

async function get(target) {
    // synchronous request
    await fetch(target, {
        method: 'GET',
        // if using "omit", nothing will happen when
        // CORS disallowed cookie
        credentials: 'include'
    }
    ).then(response => {
        if (response.ok) {
            response.json().then(json => console.log(json))
            return
        }
        throw new Error('Network response was not ok.');
    }).catch(err => console.log(err));
}

function log_cookie() {
    let cookie = document.cookie;
    console.log("Cookies: \n" + cookie)
}

// ------------- main ---------------

// cookie

function simple_request() {
    let url = server + "/cookie/create"
    get(url)
    log_cookie()
}

function delete_cookie() {
    let url = server + "/cookie/delete"
    get(url)
    log_cookie()
}

function with_httponly() {
    let url = server + "/cookie/create/httponly"
    get(url)
    log_cookie()
}

function with_secure() {
    let url = server + "/cookie/create/secure"
    get(url)
    log_cookie()
}

// ping

function ping() {
    let url = server + "/ping"
    get(url)
}

function ping_post() {
    let url = server + "/ping"
    post(url,{"msg":"hello"})
}

// login

function login() {
    let url = server + "/login"
    let username = document.getElementById("username").value
    post(url,{"username":username})
}


function auth() {
    let url = server + "/auth"
    resp = get(url)
}