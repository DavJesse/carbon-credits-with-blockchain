document.addEventListener("DOMContentLoaded", function() {
    // Add any JavaScript functionality here
    console.log("Carbon Credits Marketplace Loaded");
});

document.addEventListener("DOMContentLoaded", function() {
    const buyButtons = document.querySelectorAll('.market-item button');

    buyButtons.forEach(button => {
        button.addEventListener('click', function() {
            const item = this.parentElement.querySelector('h2').innerText;
            alert(`You have bought ${item}`);
            // Here, you can add functionality to update the user's dashboard, deduct the price, etc.
        });
    });
});
