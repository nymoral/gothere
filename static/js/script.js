function checkLen(obj, min) {
    if (obj.value.length < min) {
        obj.style.borderColor="#FF0000";
    }
    else
        obj.style.borderColor="#8A9B0F";
}

function checkEmail(obj) {
    var re = /[a-z0-9!#$%&'*+/=?^_{|}~-]+(?:\.[a-z0-9!#$%&'*+/=?^_{|}~-]+)*@(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?/;

    if (re.test(obj.value)) {
        obj.style.borderColor="#8A9B0F";
    }
    else
        obj.style.borderColor="#FF0000";
}

function checkMatch(obj) {
    s = document.getElementById("firstpassword").value;
    if (obj.value === s) {
        obj.style.borderColor="#8A9B0F";
    }
    else
        obj.style.borderColor="#FF0000";
}

function validateForm() {
    var fn = document.getElementById("first").value;
    var ln = document.getElementById("last").value;
    var em = document.getElementById("email").value;
    var ps = document.getElementById("firstpassword").value;
    var cp = document.getElementById("cpassword").value;

    if (fn.length < 1 || fn.length > 20) {
        alert("Vardas turi buti 1-20 simboliu ilgio.");
    }
    if (ln.length < 1 || ln.length > 30) {
        alert("Pavarde turi buti 1-30 simboliu ilgio.");
    }
    if (ln.length < 1 || ln.length > 30) {
        alert("Pavarde turi buti 1-30 simboliu ilgio.");
    }
    var re = /[a-z0-9!#$%&'*+/=?^_{|}~-]+(?:\.[a-z0-9!#$%&'*+/=?^_{|}~-]+)*@(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?/;
    if (! re.test(em)) {
        alert("Netinkamas el. pasto adresas.");
    }
    if (ps.length < 6) {
        alert("Slaptazodis turi buti bent 6 simboliu ilgio.");
    }
    if (ps !== cp) {
        alert("Slaptazodziai nesutampa");
    }
    return true;
}
