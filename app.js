document.addEventListener("DOMContentLoaded", function() {
    
    document.getElementById("form").addEventListener('submit', function(event) {
        event.preventDefault();
    
        const page = 3;
        const line = document.getElementById("line").value;

        const address = `http://localhost:8083/pages/${page}/${line}`
        
        axios.get(address)
            .then((response) => {
                console.log(response.data);
                const data = response.data;

                getWords(data.first_word_id, data.last_word_id)
            }).catch((err) => {
                console.error(err)
            });
    })

    document.getElementById("form-surah").addEventListener('submit', function(event) {
        event.preventDefault()

        const surah = document.getElementById("surah").value;

        const address = `http://localhost:8083/surahs/${surah}`;

        
        axios.get(address)
            .then((response) => {
                console.log(response.data);
                const data = response.data;
                document.getElementById("surah-text").innerText = data.arabic
            })
    })
})

async function getWords(first_id, last_id) {
    try {
        var text = "";
        for (let i = first_id; i <= last_id; i++) {
            let address_word = `http://localhost:8083/words/${i}`
            await axios.get(address_word)
                .then((response) => {
                    console.log(response.data);
                    text += response.data.text + " "
                }).catch((err) => {
                    console.error(err);
                });
        }

        console.log(text);
        document.getElementById('quran').innerText = text
    } catch (err) {
        console.error(err)
    }
}