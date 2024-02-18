<script setup>
    import { entity } from "../wailsjs/go/models";
import {GetBookByID, GetAllBooks, CreateBook, DeleteBook, UpdateBook, GetStudentByID, DeleteStudent, UpdateStudent, CreateStudent } from "../wailsjs/go/transport/Transport"
</script>

<template>
  <div id="sidebar">
        <div id="students">
                <button class="internal-button" @click="showFindStudentForm()">Найти ученика</button>
                <button class="internal-button" @click="showAddStudentForm()">Добавить ученика</button>
                <button class="internal-button" @click="showEditStudentForm()">Изменить ученика</button>
                <button class="internal-button" @click="showDeleteStudentForm()">Удалить ученика</button>
        </div>
        <div id="books">
                <button class="internal-button" @click="showAllBook()">Все книги</button>
                <button class="internal-button" @click="showFindBookForm()">Найти книгу</button>
                <button class="internal-button" @click="showAddBookForm()">Добавить книгу</button>
                <button class="internal-button" @click="showEditBookForm()">Изменить книгу</button>
                <button class="internal-button" @click="showDeleteBookForm()">Удалить книгу</button>    
        </div>
    </div>

    <div id="main-content">
        <div class="table"  id="studentsTable" v-show="currentForm === 'studentsTable'">
            <table>
            <thead>
                <tr>
                    <th>ID</th>
                    <th>Класс</th>
                    <th>Фамилия Имя</th>
                    <th>Книги</th>
                </tr>
            </thead>
            <tbody v-for="s in students" :key="id">
                <tr>
                    <td>{{ s.id }}</td>
                    <td>{{ s.class }}</td>
                    <td>{{ s.name }}</td>
                    <td>

                    <p v-for="b in s.books" :key="id">
                        {{ b.author }}
                        {{ b.title}}
                    </p>
                    </td>
                </tr>
            </tbody>
            </table>

        </div>

        <div class="table"  id="booksTable" v-show="currentForm === 'booksTable'">
            <table>
            <thead>
                <tr>
                    <th>ID</th>
                    <th>Автор</th>
                    <th>Название</th>
                    <th>Год издания</th>
                    <th>Количество</th>
                    <th>Обладатели</th>
                    
                </tr>
            </thead>
            <tbody v-for="book in books" :key="id">
                <tr>
                    <td>{{ book.id }}</td>
                    <td>{{ book.author }}</td>
                    <td>{{ book.title }}</td>
                    <td>{{ book.year }}</td>
                    <td>{{ book.count }}</td>
                    <td>
                        <p v-for="s in book.students" :key="id">
                            {{ s.class}}
                            {{ s.name }}
                        </p>
                    </td>
                </tr>
            </tbody> 
            </table>
        </div>

        <div class="error" v-show="currentForm === 'errorInfo'">
            <h1>
                {{ error }}
            </h1>
        </div>

        <div class="form" v-show="currentForm === 'deleteBookForm'">
            <h2>Удалить книгу</h2>
            <label for="ID" >Введите ID книги:</label>
            <input type="number" v-model="deleteBookValue">
            <button type="submit" @click="deleteBook()">Удалить</button>
        </div>

        <div class="form"  v-show="currentForm === 'findBookForm'">
            <h2>Найти книгу</h2>
            <label for="ID">Введите название или автора:</label>
            <input v-model="findBookValue" type="text" >
            <button type="submit" @click="findBook()">Найти</button>
        </div>

        <div class="form"  v-show="currentForm === 'addBookForm'">
            <h2>Добавить книгу</h2>
                <label for="Title">Название книги:</label>
                <input type="text" v-model="addBookTitle">
                <label for="Author">Автор:</label>
                <input type="text" v-model="addBookAuthor">
                <label for="count">Количество книг:</label>
                <input type="number" v-model="addBookCount">
                <label for="students">Год издания:</label>
                <input type="number" v-model="addBookYear">
                <button type="submit" @click="addBook()">Дoбавить</button>
        </div>

        <div class="form"  v-show="currentForm === 'editBookForm'">
            <h2>Изменить книгу</h2>
                <label for="ID">Введите ID книги:</label>
                <input type="number" v-model="edit_id">
                <label for="Title">Название книги:</label>
                <input type="text" v-model="edit_title">
                <label for="Author">Автор:</label>
                <input type="text" v-model="edit_author">

                <label for="count">Количество книг:</label>
                <input type="number" v-model="edit_count">
                <label for="students">Год издания:</label>
                <input type="number" v-model="edit_year">
                <button type="submit" @click="updateBook()">Изменить</button>
        </div>

        <div class="form" v-show="currentForm === 'deleteStudentForm'">
            <h2>Удалить ученика</h2>
            <label for="ID" >Введите ID ученика:</label>
            <input type="number" v-model="deleteStudentValue">
            <button type="submit" @click="deleteStudent()">Удалить</button>
        </div>

        <div class="form" v-show="currentForm === 'findStudentForm'">
            <h2>Найти ученика</h2>
            <label for="ID">Введите имя ученика:</label>
            <input v-model="findStudentValue" type="text" >
            <button type="submit" @click="findStudent()">Найти</button>
        </div>

        <div class="form" v-show="currentForm === 'addStudentForm'">
            <h2>Добавить ученика</h2>
            <label for="Title">ФИ ученика:</label>
            <input type="text" v-model="addStudentName">
            <label for="Author">Класс</label>
            <input type="text" v-model="addStudentClass">
            <button type="submit" @click="addStudent()">Дoбавить</button>
        </div>

        <div class="form" v-show="currentForm === 'editStudentForm'">
            <h2>Изменить ученика</h2>
            <label for="ID">Введите ID ученика:</label>
            <input v-model="editStudentID" type="number" >
            <label for="Name">ФИ ученика:</label>
            <input type="text" v-model="editStudentName">
            <label for="Class">Класс</label>
            <input type="text" v-model="editStudentClass">
            <button type="submit" @click="updateStudent()">Изменить</button>
        </div>
    </div>

</template>

<style scoped>
    body {
        font-family: "JetBrains Mono", sans-serif;
        margin: 0;
        padding: 0;
        display: flex;
    }

    #sidebar {
        float: left;
        height: 100%;
        width: 33%;
        background-color: #bcb9b9b3;
        padding: 20px;
        box-shadow: 2px 0px 5px #db72e4;
        border-top-right-radius: 20px;
        border-bottom-right-radius: 20px;
    }

    #main-content {
        float: right;
        margin-left: calc(67%+5px);
        width: 60%;
        right: 100px;
        flex: 1;
        padding: 20px;
    }

    #students {
        width: 90%;
        border-radius: 20px;
        background-color: #dae4e0;
        padding: 5px;
    }

    #students::before {
        height: 20px;
        color: #0c0101;
        content: "Ученики";
        display: block;
        text-align: left;
        font-size: large;
        margin-bottom: 5px;
    }
    #books {
        width: 90%;
        position: relative;
        border-radius: 20px;
        background-color:  #dae4e0;
        padding: 5px;
    }

    #books::before {
        border: #0c0101;
        color: #0c0101;
        height: 20px;
        content: "Книги";
        display: block;
        text-align: left;
        font-size: large;
        margin-bottom: 5px;
    }

    #error {
        border-radius: 20px;
        font-size: large;
        width: 90%;
    }

    .table {
        width: 100%;
        border-top-right-radius: 20px;
        border-collapse: collapse;
    }

    .table table {
        height: 100px;
        width: 100%;
        border-collapse: collapse;
    }

    .table th, tr{
        padding: 8px;
    }

    .table th {
        background-color: #0eebfb;
    }

    .table tbody tr:nth-child(even) {
        background-color: #11e9ec;
    }

    .table tbody tr:hover {
        background-color: #08f010;
    }

    .tab-button {
        background: #536dfe;
        border: none;
        color: #fff;
        padding: 10px 20px;
        display: block;
        width: 80%;
        text-align: left;
        cursor: pointer;
        border-radius: 10px;
        margin-bottom: 10px;
    }
    .form {
        background-color: #eee9e9f3;
        padding: 20px;
        border-radius: 5px;
        box-shadow: 0px 0px 10px rgb(233, 5, 249);
    }

    .form h2 {
        font-size: 24px;
        margin-bottom: 20px;
    }

    .form label {
        display: block;
        margin-bottom: 10px;
    }

    .form {
        padding: 20px;
        border: 1px solid #ccc;
        border-radius: 5px;
    }

    .internal-button {
        font-size: medium;
        height: 20%;
        position: relative;
        left: 30px;
        background: #449087e5;
        border: none;
        color: #fff;
        padding: 10px 20px;
        display: block;
        width: 80%;
        text-align: left;
        cursor: pointer;
        border-radius: 10px;
        margin-bottom: 10px;
    }

    .tab-button:hover {
        background-color: #4255ff;
    }
    .form {
        color: #0c0101;
        margin-top: 20px;
    }
    .form label {
        display: block;
        margin-top: 10px;
    }
    .form input {
        width: 100%;
        padding: 10px;
        margin-top: 5px;
        border: 1px solid #ccc;
        border-radius: 5px;
    }
    .form button {
        background-color: #536dfe;
        color: #fff;
        padding: 10px 20px;
        border: none;
        border-radius: 5px;
        cursor: pointer;
        margin-top: 10px;
    }
    .form button:hover {
        background-color: #304ffe;
    }
</style>

<script>

export default {
  data() {
    return {
      error: '',
      books: [],
      students: [],
      currentForm: ''
    };
  },
  methods: {
    findBook: async function() {
            this.books = []
            const word = this.findBookValue;
            const book = await GetBookByID(word)

            console.log(book)
            this.books.push(book);


            this.currentForm = 'booksTable';
    },

    allBook: async function() {
        this.books = await GetAllBooks();
        this.currentForm = 'booksTable';
    },

    deleteBook: async function() {
        const id = this.deleteBookValue;
        this.error = await DeleteBook(id);
        this.currentForm = 'errorInfo'
    },

    updateBook: async function() {
        const form = {
            id: this.edit_id,
            title: this.edit_title,
            author: this.edit_author,
            count: this.edit_count,
            year: this.edit_year
        };

        let book = new entity.Book(form);

        this.error = await UpdateBook(book);
        this.currentForm = 'errorInfo'
    },

    addBook: async function() {
        const form = {
            title: this.addBookTitle,
            author: this.addBookAuthor,
            count: this.addBookCount,
            year: this.addBookYear
        };

        let book = new entity.Book(form);

        this.error = await CreateBook(book);

        this.currentForm = 'errorInfo'
    },

    findStudent: async function() {
        this.students = []
        const name = this.findStidentValue;
        const student = await GetStudentByID(name)

        this.students.push(student);
        this.currentForm = 'studentsTable';
    },

    allStudent: async function() {
        this.books = await GetAllBooks();
        this.currentForm = 'booksTable';
    },

    deleteStudent: async function() {
        const id = this.deleteBookValue;
        this.error = await DeleteStudent(id);
        this.currentForm = 'errorInfo'
    },

    updateStudent: async function() {
        const form = {
            id: this.editStudentID,
            class: this.editStudentClass,
            name: this.addStudentName
        };

        let student = new entity.Student(form);

        this.error = await UpdateStudent(student);
        this.currentForm = 'errorInfo'
    },

    addStudent: async function() {
        const form = {
            class: this.addStudentClass,
            name: this.addStudentName
        };

        let student = new entity.Student(form);

        this.error = await CreateStudent(student);
        this.currentForm = 'errorInfo'
    },

    showAllBook() {
        this.currentForm = 'allBooks'
        this.allBook()
    },
    showDeleteBookForm() {
      this.currentForm = 'deleteBookForm';
    },
    showFindBookForm() {
      this.currentForm = 'findBookForm';
    },
    showAddBookForm() {
      this.currentForm = 'addBookForm';
    },
    showEditBookForm() {
      this.currentForm = 'editBookForm';
    },


    showFindStudentForm() {
        this.currentForm = 'findStudentForm'
    },
    showAddStudentForm() {
        this.currentForm = 'addStudentForm'
    },
    showEditStudentForm() {
        this.currentForm = 'editStudentForm'
    },
    showDeleteStudentForm() {
        this.currentForm = 'deleteStudentForm'
    }
  }
};
</script>