document.addEventListener("DOMContentLoaded", function () {
    let allExperienceData = [];
    let currentIndex = 0;
    const experienceContainer = document.getElementById('experienceContainer');
    const loadMoreButton = document.getElementById('loadMoreButton');
    const experiencesPerPage = 2; // 한 번에 표시할 경험의 수

    function displayExperience() {
        const endIndex = Math.min(currentIndex + experiencesPerPage, allExperienceData.length);

        for (let i = currentIndex; i < endIndex; i++) {
            const experience = allExperienceData[i];

            const experienceCard = ` 
                <div class="card shadow border-0 rounded-4 mb-5" data-id="${experience.ID}">
                    <div class="card-body p-5">
                        <div class="row align-items-center gx-5">
                            <div class="col text-center text-lg-start mb-4 mb-lg-0">
                                <div class="bg-light p-4 rounded-4">
                                    <div class="date-container">
                                        <div class="text-primary fw-bolder">${experience.period}</div>
                                    </div>
                                    <div class="text-container">
                                        <div class="small fw-bolder">${experience.role}</div>
                                        <div class="small text-muted">${experience.company}</div>
                                    </div>
                                </div>
                            </div>
                            <div class="col-lg-8">
                                <div>${experience.description}</div>
                            </div>
                        </div>
                    </div>
                    <button class="btn btn-light position-absolute top-0 start-0 m-2 p-1 edit-button" 
                    style="border-radius: 50%; z-index: 10;">
                        <i class="bi bi-pencil"></i>
                    </button>
                </div>
            `;

            experienceContainer.insertAdjacentHTML('beforeend', experienceCard);
        }

        currentIndex = endIndex;

        if (currentIndex >= allExperienceData.length) {
            loadMoreButton.style.display = 'none';
        }
    }

    fetch('/returnresume')
        .then(response => response.json())
        .then(data => {
            allExperienceData = data.exps;
            displayExperience();
        })
        .catch(error => console.error('Error fetching data:', error));

    loadMoreButton.addEventListener('click', function () {
        displayExperience();
    });

    document.getElementById('upload').addEventListener('click', function () {
        const uniqueUploadId = `uploadCard-${Date.now()}`;
        const newCard = `
            <div class="card shadow border-0 rounded-4 mb-5" id="${uniqueUploadId}">
                <div class="card-body p-5">
                    <h3 class="fw-bolder mb-4">Add New Experience</h3>
                    <form id="uploadForm-${uniqueUploadId}">
                        <div class="mb-3">
                            <label for="periodStart-${uniqueUploadId}" class="form-label">Start Date</label>
                            <input type="month" class="form-control" id="periodStart-${uniqueUploadId}" required>
                        </div>
                        <div class="mb-3">
                            <label for="periodEnd-${uniqueUploadId}" class="form-label">End Date</label>
                            <input type="month" class="form-control" id="periodEnd-${uniqueUploadId}" required>
                        </div>
                        <div class="mb-3">
                            <label for="role-${uniqueUploadId}" class="form-label">Role</label>
                            <input type="text" class="form-control" id="role-${uniqueUploadId}" required>
                        </div>
                        <div class="mb-3">
                            <label for="company-${uniqueUploadId}" class="form-label">Company</label>
                            <input type="text" class="form-control" id="company-${uniqueUploadId}" required>
                        </div>
                        <div class="mb-3">
                            <label for="description-${uniqueUploadId}" class="form-label">Description</label>
                            <textarea class="form-control" id="description-${uniqueUploadId}" rows="3" required></textarea>
                        </div>
                        <button type="submit" class="btn btn-primary">Submit</button>
                        <button type="button" id="cancelUpload-${uniqueUploadId}" class="btn btn-secondary">Cancel</button>
                    </form>
                </div>
            </div>
        `;

        experienceContainer.insertAdjacentHTML('afterbegin', newCard);

        document.getElementById(`uploadForm-${uniqueUploadId}`).addEventListener('submit', function (event) {
            event.preventDefault();

            const periodStart = document.getElementById(`periodStart-${uniqueUploadId}`).value;
            const periodEnd = document.getElementById(`periodEnd-${uniqueUploadId}`).value;
            const role = document.getElementById(`role-${uniqueUploadId}`).value;
            const company = document.getElementById(`company-${uniqueUploadId}`).value;
            const description = document.getElementById(`description-${uniqueUploadId}`).value;

            const formatPeriod = (start, end) => {
                const [startYear, startMonth] = start.split("-");
                const [endYear, endMonth] = end.split("-");
                return `${startYear}년 ${startMonth}월 ~ ${endYear}년 ${endMonth}월`;
            };

            const periodFormatted = formatPeriod(periodStart, periodEnd);

            fetch('/uploadresume', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({
                    period: periodFormatted,
                    role,
                    company,
                    description
                })
            })
            .then(response => response.json())
            .then(result => {
                alert('Experience added successfully!');
                document.getElementById(uniqueUploadId).remove();
            })
            .catch(error => console.error('Error:', error));
        });

        document.getElementById(`cancelUpload-${uniqueUploadId}`).addEventListener('click', function () {
            document.getElementById(uniqueUploadId).remove();
        });
    });

    experienceContainer.addEventListener('click', function (event) {
        if (event.target.closest('.edit-button')) {
            const card = event.target.closest('.card');
            const cardId = card.getAttribute('data-id');

            const periodText = card.querySelector('.text-primary').textContent;
            const [startPeriod, endPeriod] = periodText.split(' ~ ').map(period => {
                const [year, month] = period.trim().split('년 ');
                return `${year.trim()}-${month.replace('월', '').trim().padStart(2, '0')}`;
            });

            const editCardId = `editCard-${Date.now()}`;
            const editCard = `
                <div class="card shadow border-0 rounded-4 mb-5" id="${editCardId}">
                    <div class="card-body p-5">
                        <h3 class="fw-bolder mb-4">Edit Experience</h3>
                        <form id="editForm-${editCardId}">
                            <input type="hidden" id="editId" value="${cardId}">
                            <div class="mb-3">
                                <label for="editPeriodStart-${editCardId}" class="form-label">Start Date</label>
                                <input type="month" class="form-control" id="editPeriodStart-${editCardId}" value="${startPeriod}" required>
                            </div>
                            <div class="mb-3">
                                <label for="editPeriodEnd-${editCardId}" class="form-label">End Date</label>
                                <input type="month" class="form-control" id="editPeriodEnd-${editCardId}" value="${endPeriod}" required>
                            </div>
                            <div class="mb-3">
                                <label for="editRole-${editCardId}" class="form-label">Role</label>
                                <input type="text" class="form-control" id="editRole-${editCardId}" value="${card.querySelector('.small.fw-bolder').textContent}" required>
                            </div>
                            <div class="mb-3">
                                <label for="editCompany-${editCardId}" class="form-label">Company</label>
                                <input type="text" class="form-control" id="editCompany-${editCardId}" value="${card.querySelector('.small.text-muted').textContent}" required>
                            </div>
                            <div class="mb-3">
                                <label for="editDescription-${editCardId}" class="form-label">Description</label>
                                <textarea class="form-control" id="editDescription-${editCardId}" rows="3" required>${card.querySelector('.col-lg-8 div').textContent}</textarea>
                            </div>
                            <button type="submit" class="btn btn-primary">Save Changes</button>
                            <button type="button" id="cancelEdit-${editCardId}" class="btn btn-secondary">Cancel</button>
                        </form>
                    </div>
                </div>
            `;

            card.insertAdjacentHTML('beforeend', editCard);

            document.getElementById(`editForm-${editCardId}`).addEventListener('submit', function (event) {
                event.preventDefault();

                const periodStart = document.getElementById(`editPeriodStart-${editCardId}`).value;
                const periodEnd = document.getElementById(`editPeriodEnd-${editCardId}`).value;
                const role = document.getElementById(`editRole-${editCardId}`).value;
                const company = document.getElementById(`editCompany-${editCardId}`).value;
                const description = document.getElementById(`editDescription-${editCardId}`).value;

                const formatPeriod = (start, end) => {
                    const [startYear, startMonth] = start.split("-");
                    const [endYear, endMonth] = end.split("-");
                    return `${startYear}년 ${startMonth}월 ~ ${endYear}년 ${endMonth}월`;
                };

                const periodFormatted = formatPeriod(periodStart, periodEnd);

                fetch('/editresume', {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json'
                    },
                    body: JSON.stringify({
                        id: cardId,
                        period: periodFormatted,
                        role,
                        company,
                        description
                    })
                })
                .then(response => response.json())
                .then(result => {
                    alert('Experience updated successfully!');
                    document.getElementById(editCardId).remove();
                    card.querySelector('.text-primary').textContent = periodFormatted;
                    card.querySelector('.small.fw-bolder').textContent = role;
                    card.querySelector('.small.text-muted').textContent = company;
                    card.querySelector('.col-lg-8 div').textContent = description;
                })
                .catch(error => console.error('Error:', error));
            });

            document.getElementById(`cancelEdit-${editCardId}`).addEventListener('click', function () {
                document.getElementById(editCardId).remove();
            });
        }
    });
});
