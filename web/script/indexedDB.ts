class Database {
    private db: Maybe<IDBDatabase> = null;

    constructor(
        public name: string,
        public version: number,
    ) {}

    public open() {
        const request = indexedDB.open(this.name, this.version);

        request.addEventListener("success", () => {
            this.db = request.result;
        });
    }

    public conn(callback: (db: IDBDatabase) => void) {
        if (this.db !== null) {
            callback(this.db);
        }
    }
}

const db = new Database("sample-database", 1);

const request = indexedDB.open("sample-database", 1);

request.addEventListener("upgradeneeded", function () {
    this.result.addEventListener("abort", function () {
        console.error("Load: abort");
    });

    this.result.addEventListener("close", function () {
        console.error("Load: close");
    });

    this.result.addEventListener("versionchange", function () {
        console.error("Load: version change");
    });

    const user = this.result.createObjectStore("user", {
        keyPath: "id",
        autoIncrement: true,
    });

    user.add({ id: 1, firstName: "John", lastName: "Doe" });
    user.add({ id: 2, firstName: "Sam", lastName: "Smith" });
    user.add({ id: 3, firstName: "Veronica", lastName: "Sullivan" });
    user.add({ id: 4, firstName: "Hannah", lastName: "Mulligan" });

    const insert = () => {
        const newUser = {
            firstName: "Sally",
            lastName: "Doe",
        };

        const transaction = this.result.transaction(["user"], "readwrite");
        const store = transaction.objectStore("user");
        const query = store.add(newUser);

        query.addEventListener("success", function () {
            console.log("Success", newUser);
        });

        transaction.addEventListener("complete", function () {
            console.log("Complete", newUser);
        });

        transaction.addEventListener("error", function () {
            console.log("Transaction error");
        });
    };
});

request.addEventListener("success", function () {
    console.log("EventListener: Success");
});

request.addEventListener("error", function () {
    console.log("EventListener: Error");
});

request.addEventListener("blocked", function () {
    console.log("EventListener: Blocked");
});
