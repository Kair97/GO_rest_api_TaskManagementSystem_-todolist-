const priorities = document.querySelectorAll(".priorityList li");

priorities.forEach(item => {
    item.addEventListener("click", () => {
        priorities.forEach(p => p.classList.remove("active", "orange", "green", "red"));
        const value = Number(item.dataset.value);
        if (value == 5){
            item.classList.add("red");
        }else if (value == 3 || value == 4){
            item.classList.add("orange");
        }else{
            item.classList.add("green");
        }
    });
});
console.log("hi")

const filterButtons = document.querySelectorAll(".filter-btn");
let currentFilter = "all";

filterButtons.forEach(button => {
    button.addEventListener("click", ()=>{
        filterButtons.forEach(b => b.classList.remove("active"));
        button.classList.add("active");
        currentFilter = button.dataset.filter;
        loadTasks();
    });
});

const API_URL = "/tasks";

async function createTask() {
    const title = document.getElementById("titleInput");
    const description = document.getElementById("descInput");

    const titleValue = title.value.trim();
    if (!titleValue){
        alert("Title is needed");
        return;
    }
    const descriptionValue = description.value.trim();

    const activePriority = document.querySelector(
        ".priorityList li.green, .priorityList li.orange, .priorityList li.red"
    );

    if (!activePriority){
        alert("Select priority please");
        return;
    }

    const priority = Number(activePriority.dataset.value);

    try {
        const response = await fetch(API_URL, {
            method: "POST",
            headers: {"Content-Type": "application/json"},
            body: JSON.stringify({
                title: titleValue,
                description: descriptionValue || "",
                priority: priority,
                completed: false
            })
        });

        if (!response.ok) {
            alert("Failed to create task");
            return;
        }

        title.value = "";
        description.value = "";
        priorities.forEach(p => 
            p.classList.remove("active", "orange", "green", "red")
        );

        loadTasks();
    } catch (error) {
        console.error("Error creating task:", error);
        alert("Error creating task");
    }
}

async function loadTasks(){
    try {
        const res = await fetch(API_URL);
        if (!res.ok) {
            alert("Server error");
            return;
        }
        const tasks = await res.json();
        
        updateStatistics(tasks);

        let filteredTasks = tasks;

        if (currentFilter === "active"){
            filteredTasks = tasks.filter(t => !t.completed);
        } else if (currentFilter === "completed"){
            filteredTasks = tasks.filter(t => t.completed);
        }

        const tasksDiv = document.getElementById("tasksList");
        tasksDiv.innerHTML = "";

        filteredTasks.forEach(task => {
            const div = document.createElement("div");
            div.className = "task-card mb-3";

            div.innerHTML = `
                <div class="task-left">
                    <input type="checkbox" ${task.completed ? "checked" : ""}>
                    <div>
                        <h6 class="task-title">${escapeHtml(task.title)}</h6>
                        ${task.description ? `<p class="task-desc">${escapeHtml(task.description)}</p>` : ""}
                        <div class="prior-div">
                            <p>Priority ${task.priority}</p>
                        </div>
                    </div>
                </div>

                <div class="task-right">
                    <button class="btn btn-sm edit-task">Edit</button>
                    <button class="btn btn-sm btn-outline-danger delete-btn">Delete</button>
                </div>
            `;

            tasksDiv.appendChild(div);
            
            const editBtn = div.querySelector(".edit-task");
            editBtn.addEventListener("click", ()=>{
                openUpdateModal(
                    task.id,
                    task.title,
                    task.description,
                    task.completed,
                    task.priority
                );
            });

            const deleteBtn = div.querySelector(".delete-btn");
            deleteBtn.addEventListener("click", ()=>{
                deleteTask(task.id);
            });

            const checkbox = div.querySelector("input[type='checkbox']");
            checkbox.addEventListener("change", async()=>{
                try {
                    const res = await fetch(`/tasks/${task.id}`, {
                        method: "PATCH",
                        headers: {"Content-Type": "application/json"},
                        body: JSON.stringify({
                            completed: checkbox.checked
                        })
                    });
                    
                    if (!res.ok) {
                        alert("Failed to update task");
                        checkbox.checked = !checkbox.checked;
                        return;
                    }
                    
                    loadTasks();
                } catch (error) {
                    console.error("Error updating task:", error);
                    alert("Failed to update task");
                    checkbox.checked = !checkbox.checked;
                }
            });
        });
    } catch (error) {
        console.error("Error loading tasks:", error);
        alert("Error loading tasks");
    }
}

function updateStatistics(tasks) {
    const totalTasks = tasks.length;
    const activeTasks = tasks.filter(t => !t.completed).length;
    const completedTasks = tasks.filter(t => t.completed).length;

    document.getElementById("totalTasks").textContent = totalTasks;
    document.getElementById("activeTasks").textContent = activeTasks;
    document.getElementById("completedTasks").textContent = completedTasks;
}

function escapeHtml(text) {
    const div = document.createElement('div');
    div.textContent = text;
    return div.innerHTML;
}

let currentTaskId = null;

function openUpdateModal(id, title, description, completed, priority){
    currentTaskId = id;

    document.getElementById("updateTitle").value = title || "";
    document.getElementById("updateDescription").value = description || "";
    document.getElementById("updateCompleted").checked = completed;
    document.getElementById("updatePriority").value = priority;

    const modal = new bootstrap.Modal(
        document.getElementById("updateModal")
    );

    modal.show();
}

async function submitUpdate(){
    if (!currentTaskId) return;
    
    const title = document.getElementById("updateTitle").value.trim();
    const description = document.getElementById("updateDescription").value.trim();
    const completed = document.getElementById("updateCompleted").checked;
    const priority = Number(document.getElementById("updatePriority").value);

    if (!title) {
        alert("Title is required");
        return;
    }

    try {
        const response = await fetch(`/tasks/${currentTaskId}`, {
            method: "PATCH",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify({
                title, 
                description: description || "",
                completed,
                priority 
            })
        });

        if (!response.ok) {
            alert("Failed to update task");
            return;
        }

        currentTaskId = null;

        const modalEl = document.getElementById("updateModal");
        const modalInstance = bootstrap.Modal.getInstance(modalEl);
        if (modalInstance) {
            modalInstance.hide();
        }

        loadTasks();
    } catch (error) {
        console.error("Error updating task:", error);
        alert("Error updating task");
    }
}

async function deleteTask(id){
    if (!confirm("Are you sure you want to delete this task?")) {
        return;
    }

    try {
        const response = await fetch(`${API_URL}/${id}`, {
            method: "DELETE"
        });

        if (!response.ok) {
            alert("Failed to delete task");
            return;
        }

        loadTasks();
    } catch (error) {
        console.error("Error deleting task:", error);
        alert("Error deleting task");
    }
}

loadTasks();