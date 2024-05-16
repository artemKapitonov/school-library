<script setup>
    import { entity } from "../wailsjs/go/models";
import {GetAllBooks, RegisterStudentBook, CreateBook, DeleteBook, UpdateBook, SearchBook, DeleteStudent, UpdateStudent, CreateStudent, SearchStudent, GetAllStudents } from "../wailsjs/go/transport/Transport"
</script>

<template>
    <div id="sidebar">
        <div id="students">
                <button class="internal-button" @click="showAllStudents()">Все ученики</button>
                <button class="internal-button" @click="showFindStudentForm()">Найти ученика</button>
                <button class="internal-button" @click="showAddStudentForm()">Добавить ученика</button>
                <button class="internal-button" @click="showEditStudentForm()">Изменить ученика</button>
                <button class="internal-button" @click="showDeleteStudentForm()">Удалить ученика</button>
        </div>
        <div>
            <br>
        </div>
        <div id="books">
                <button class="internal-button" @click="showAllBook()">Все книги</button>
                <button class="internal-button" @click="showFindBookForm()">Найти книгу</button>
                <button class="internal-button" @click="showAddBookForm()">Добавить книгу</button>
                <button class="internal-button" @click="showEditBookForm()">Изменить книгу</button>
                <button class="internal-button" @click="showDeleteBookForm()">Удалить книгу</button>    
        </div>

        <div id="register">
            <button class="internal-button" @click="showRegisterBookForStudent()">Записать книгу ученику</button>
            <button class="internal-button" @click="showUnregisterBookForStudent()">Снять запись книги</button>
        </div>
    </div>

    <div id="main-content">

        <div class="form" v-show="currentForm === 'studentsTable'">
            <div class="table"  id="studentsTable" >
                <table>
                <thead>
                    <tr>
                        <th  class="br1">ID</th>
                        <th>Класс</th>
                        <th>Фамилия Имя</th>
                        <th  class="br2">Книги</th>
                    </tr>
                </thead>
                <tbody v-for="s in students" :key="id">
                    <tr>
                        <td>{{ s.id }}</td>
                        <td>{{ s.class }}</td>
                        <td>{{ s.name }}</td>
                        <td>

                        <p v-for="b in s.books" :key="id">
                            {{ b.id }}
                            {{ b.author }}
                            {{ b.title}}
                        </p>
                        </td>
                    </tr>
                </tbody>
                </table>
            </div>
        </div>

        <div class="form" v-show="currentForm === 'booksTable'">
            <div class="table"  id="booksTable" >
                <table>
                <thead>
                    <tr>
                        <th class="br1">ID</th>
                        <th>Автор</th>
                        <th>Название</th>
                        <th>Год издания</th>
                        <th>Количество</th>
                        <th class="br2">Обладатели</th>
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
                                {{ s.id }}
                                {{ s.class}}
                                {{ s.name }}
                            </p>
                        </td>
                    </tr>
                </tbody> 
                </table>
            </div>
        </div>

        

        <div class="form" v-show="currentForm === 'errorInfo'">
            <h1>
                {{ error }}
            </h1>
        </div>

        <div class="form" v-show="currentForm === 'registerForm'">
            <h2>Записать книгу ученику</h2>
            <label for="ID" >Введите ID ученика:</label>
            <input type="number" v-model="registerStudentValue">
            <label for="ID" >Введите ID книги:</label>
            <input type="number" v-model="registerBookValue">
            <button type="submit" @click="RegisterBookForStudent()">Записать</button>
        </div>

        <div class="form" v-show="currentForm === 'unregisterForm'">
            <h2>Снять запись книги</h2>
            <label for="ID" >Введите ID ученика:</label>
            <input type="number" v-model="registerStudentValue">
            <label for="ID" >Введите ID книги:</label>
            <input type="number" v-model="registerBookValue">
            <button type="submit" @click="RegisterBookForStudent()">Снять</button>
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
        background-color: rgba(229, 227, 227, 0.7);
        padding: 20px;
        box-shadow: 0px 0px 10px rgb(14, 117, 233);
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
        margin: 10px;
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

    #register {
        width: 90%;
        position: relative;
        border-radius: 20px;
        background-color:  #dae4e0;
        padding: 5px;
        margin: 10px;
    }


    #register::before {
        border: #0c0101;
        color: #0c0101;
        height: 20px;
        content: "Записи";
        display: block;
        text-align: left;
        font-size: large;
        margin-bottom: 5px;
    }
    


    .table {
        width: 100%; 
    }

    .table table {
        height: 100px;
        width: 100%;
        border-collapse: collapse;
    }

    .br1{
          border-radius: 20px 0 0 0;
    }
    .br2{
        border-radius: 0 20px 0 0;
    }

    .table th, tr{
        padding: 8px;
    }

    .table th {
        background-color: #ffffff;
    }

    .table tbody tr:nth-child(even) {
        background-color: #ec1111;
    }
​

    .table tbody tr:hover {
        background-color: rgba(133, 133, 174, 0.77);
    }
    
    .form {
        background-color: #e3e3e3f3;
        padding: 20px;
        border-radius: 5px;
        box-shadow: 0px 0px 10px rgb(14, 117, 233);
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
        background: #3f85a3;
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
        background-color: #4b64f5;
        color: #fff;
        padding: 10px 20px;
        border: none;
        border-radius: 5px;
        cursor: pointer;
        margin-top: 10px;
    }
    .form button:hover {
        background-color: #3f85a3;
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
            this.books = await SearchBook(word)

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
        const name = this.findStudentValue;

        this.students = await SearchStudent(name)
        this.currentForm = 'studentsTable';
    },

    allStudent: async function() {
        this.students = []
        this.students = await GetAllStudents();
        this.currentForm = 'studentsTable';
    },

    deleteStudent: async function() {
        const id = this.deleteStudentValue;
        this.error = await DeleteStudent(id);
        this.currentForm = 'errorInfo'
    },

    updateStudent: async function() {
        const form = {
            id: this.editStudentID,
            class: this.editStudentClass,
            name: this.editStudentName
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

    RegisterBookForStudent: async function() {
        let book_id = this.registerBookValue
        let student_id = this.registerStudentValue

        this.error = await RegisterStudentBook(book_id, student_id)
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



    showAllStudents() {
        this.allStudent()
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
    }, 


    showRegisterBookForStudent() {
        this.currentForm = 'registerForm'
    },

    showUnregisterBookForStudent() {
        this.currentForm = 'unregisterForm'
    }
  }
};
</script>