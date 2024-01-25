const request = indexedDB.open("sample-database", 1);

request.addEventListener("upgradeneeded", function () {
    const store = this.result.createObjectStore("user", {
        keyPath: "id",
    });

    store.add({ id: 1, firstName: "John", lastName: "Doe" });
    store.add({ id: 2, firstName: "John", lastName: "Doe" });
    store.add({ id: 3, firstName: "John", lastName: "Doe" });
    store.add({ id: 4, firstName: "John", lastName: "Doe" });
});
